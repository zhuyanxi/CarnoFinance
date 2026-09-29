package main

import (
	"context"
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

type DailyBasisRecord struct {
	TradeDate   string  `json:"trade_date"`
	Contract    string  `json:"contract"`
	ExpiryDate  string  `json:"expiry_date"`
	SpotClose   float64 `json:"spot_close"`
	FutureClose float64 `json:"future_close"`
	BasisMetrics
}

type klineResponse struct {
	RC   int `json:"rc"`
	Data *struct {
		Name   string   `json:"name"`
		Klines []string `json:"klines"`
	} `json:"data"`
}

func (c *EastmoneyClient) DailyHistory(ctx context.Context, calendarDays int) ([]DailyBasisRecord, error) {
	if calendarDays <= 0 {
		return nil, fmt.Errorf("history days must be positive")
	}
	contracts, err := c.fetchIMContracts(ctx)
	if err != nil {
		return nil, fmt.Errorf("fetch IM contracts: %w", err)
	}
	today := time.Now().In(chinaTimeZone)
	startDate := today.AddDate(0, 0, -calendarDays).Format("20060102")
	endDate := today.Format("20060102")
	spotCloses, err := c.fetchDailyCloses(ctx, csi1000SecurityID, startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("fetch CSI 1000 daily prices: %w", err)
	}

	rows := make([]DailyBasisRecord, 0, len(contracts)*calendarDays/2)
	for _, contract := range contracts {
		futureCloses, err := c.fetchDailyCloses(ctx, fmt.Sprintf("%d.%s", contract.Market, contract.Code), startDate, endDate)
		if err != nil {
			return nil, fmt.Errorf("fetch %s daily prices: %w", contract.Name, err)
		}
		for tradeDateText, futureClose := range futureCloses {
			spotClose, exists := spotCloses[tradeDateText]
			if !exists {
				continue
			}
			tradeDate, err := time.Parse(time.DateOnly, tradeDateText)
			if err != nil {
				return nil, fmt.Errorf("parse trade date %q: %w", tradeDateText, err)
			}
			metrics, err := calculateBasis(spotClose, futureClose, tradeDate, contract.ExpiryDate)
			if err != nil {
				return nil, fmt.Errorf("calculate daily basis for %s on %s: %w", contract.Name, tradeDateText, err)
			}
			rows = append(rows, DailyBasisRecord{
				TradeDate:    tradeDateText,
				Contract:     contract.Name,
				ExpiryDate:   contract.ExpiryDate.Format(time.DateOnly),
				SpotClose:    spotClose,
				FutureClose:  futureClose,
				BasisMetrics: metrics,
			})
		}
	}
	if len(rows) == 0 {
		return nil, fmt.Errorf("no overlapping CSI 1000 and IM daily bars found")
	}
	sort.Slice(rows, func(left, right int) bool {
		if rows[left].TradeDate == rows[right].TradeDate {
			return rows[left].Contract < rows[right].Contract
		}
		return rows[left].TradeDate < rows[right].TradeDate
	})
	return rows, nil
}

func (c *EastmoneyClient) fetchDailyCloses(ctx context.Context, securityID, startDate, endDate string) (map[string]float64, error) {
	query := url.Values{
		"secid":   {securityID},
		"fields1": {"f1,f2,f3,f4,f5,f6"},
		"fields2": {"f51,f52,f53,f54,f55,f56,f57"},
		"klt":     {"101"},
		"fqt":     {"0"},
		"beg":     {startDate},
		"end":     {endDate},
	}
	var payload klineResponse
	if err := c.getJSON(ctx, eastmoneyKlineURL, query, &payload); err != nil {
		return nil, err
	}
	if payload.RC != 0 || payload.Data == nil {
		return nil, fmt.Errorf("Eastmoney returned no daily bars for %s (rc=%d)", securityID, payload.RC)
	}
	return parseDailyKlines(payload.Data.Klines)
}

func parseDailyKlines(klines []string) (map[string]float64, error) {
	closes := make(map[string]float64, len(klines))
	for _, line := range klines {
		parts := strings.Split(line, ",")
		if len(parts) < 3 {
			return nil, fmt.Errorf("unexpected Eastmoney kline format %q", line)
		}
		if _, err := time.Parse(time.DateOnly, parts[0]); err != nil {
			return nil, fmt.Errorf("invalid trade date in Eastmoney kline %q: %w", line, err)
		}
		closePrice, err := strconv.ParseFloat(parts[2], 64)
		if err != nil || closePrice <= 0 {
			return nil, fmt.Errorf("invalid close price in Eastmoney kline %q", line)
		}
		closes[parts[0]] = closePrice
	}
	if len(closes) == 0 {
		return nil, fmt.Errorf("Eastmoney returned an empty daily kline series")
	}
	return closes, nil
}