package eastmoney

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"time"
)

const (
	listURLTemplate  = "https://push2.eastmoney.com/api/qt/clist/get"
	defaultPageSize  = 100
	defaultUserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/136.0.0.0 Safari/537.36"
	ashareFS         = "m:0+t:6,m:0+t:80,m:1+t:2,m:1+t:23"
	bjFS             = "m:0+t:81+s:2048"
	quoteFields      = "f12,f13,f14"
	utToken          = "bd1d9ddb04089700cf9c27f6f7426281"
)

type Client struct {
	hc *http.Client
}

type Stock struct {
	Code       string
	Name       string
	Exchange   string
	MarketType string
}

type stockListResponse struct {
	Data stockListData `json:"data"`
}

type stockListData struct {
	Total int             `json:"total"`
	Diff  []stockListItem `json:"diff"`
}

type stockListItem struct {
	Code   string `json:"f12"`
	Market int    `json:"f13"`
	Name   string `json:"f14"`
}

func New() *Client {
	return &Client{
		hc: &http.Client{Timeout: 20 * time.Second},
	}
}

func (c *Client) GetAShareList(ctx context.Context) ([]Stock, error) {
	ashares, err := c.fetchList(ctx, ashareFS, false)
	if err != nil {
		return nil, err
	}
	bjShares, err := c.fetchList(ctx, bjFS, true)
	if err != nil {
		return nil, err
	}

	byCode := make(map[string]Stock, len(ashares)+len(bjShares))
	for _, item := range ashares {
		byCode[item.Exchange+item.Code] = item
	}
	for _, item := range bjShares {
		byCode[item.Exchange+item.Code] = item
	}

	stocks := make([]Stock, 0, len(byCode))
	for _, item := range byCode {
		stocks = append(stocks, item)
	}
	sort.Slice(stocks, func(left, right int) bool {
		if stocks[left].Exchange == stocks[right].Exchange {
			return stocks[left].Code < stocks[right].Code
		}
		return stocks[left].Exchange < stocks[right].Exchange
	})
	return stocks, nil
}

func (c *Client) fetchList(ctx context.Context, fs string, forceBJ bool) ([]Stock, error) {
	page := 1
	totalPages := 1
	stocks := make([]Stock, 0)

	for {
		requestURL := buildListURL(page, defaultPageSize, fs)
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
		if err != nil {
			return nil, err
		}
		req.Header.Set("User-Agent", defaultUserAgent)
		req.Header.Set("Accept", "application/json, text/plain, */*")

		resp, err := c.hc.Do(req)
		if err != nil {
			return nil, err
		}

		var payload stockListResponse
		decodeErr := json.NewDecoder(resp.Body).Decode(&payload)
		resp.Body.Close()
		if decodeErr != nil {
			return nil, fmt.Errorf("decode eastmoney stock list: %w", decodeErr)
		}

		if page == 1 && payload.Data.Total > 0 {
			totalPages = (payload.Data.Total + defaultPageSize - 1) / defaultPageSize
		}

		for _, item := range payload.Data.Diff {
			exchange := exchangeForItem(item, forceBJ)
			if exchange == "" {
				continue
			}
			stocks = append(stocks, Stock{
				Code:       item.Code,
				Name:       item.Name,
				Exchange:   exchange,
				MarketType: exchange,
			})
		}

		if page >= totalPages || len(payload.Data.Diff) == 0 {
			break
		}
		page++
	}

	return stocks, nil
}

func buildListURL(page, pageSize int, fs string) string {
	query := make(url.Values)
	query.Set("pn", strconv.Itoa(page))
	query.Set("pz", strconv.Itoa(pageSize))
	query.Set("po", "1")
	query.Set("np", "1")
	query.Set("ut", utToken)
	query.Set("fltt", "2")
	query.Set("invt", "2")
	query.Set("fid", "f12")
	query.Set("fs", fs)
	query.Set("fields", quoteFields)
	return listURLTemplate + "?" + query.Encode()
}

func exchangeForItem(item stockListItem, forceBJ bool) string {
	if forceBJ {
		return "BJ"
	}

	switch item.Market {
	case 1:
		return "SH"
	case 0:
		return "SZ"
	default:
		return ""
	}
}
