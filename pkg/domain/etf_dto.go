package domain

import (
	"sort"

	"github.com/samber/lo"
	"github.com/zhuyanxi/CarnoFinance/pkg/helper"
)

type RSRSDto struct {
	TsCode string  `json:"ts_code,omitempty"`
	Name   string  `json:"name,omitempty"`
	RSRS   float64 `json:"rsrs,omitempty"`
}

func (d *Domain) GetETFRSRSList(period int, order, date string) ([]RSRSDto, error) {
	etfList, err := d.GetETFCodeList()
	if err != nil {
		return nil, err
	}

	var rets []RSRSDto
	for _, etf := range etfList {
		var pricesList []float64
		err := d.db.NewSelect().Model((*ETFDailyPrice)(nil)).Column("close").
			Where("ts_code=?", etf.TSCode).Where("trade_date<=?", date).
			Order("trade_date DESC").
			Limit(period).
			Scan(d.ctx, &pricesList)
		if err != nil {
			return nil, err
		}

		reversePrice := lo.Reverse(pricesList)
		score := helper.RSRS(reversePrice)
		rets = append(rets, RSRSDto{
			TsCode: etf.TSCode,
			Name:   etf.Name,
			RSRS:   score,
		})
	}

	sort.Slice(rets, func(i, j int) bool {
		if order == "asc" {
			return rets[i].RSRS < rets[j].RSRS // sort asc
		}
		return rets[i].RSRS > rets[j].RSRS // sort desc
	})

	return rets, nil
}

func (d *Domain) GetETFClosePrices(code string, count int) ([]string, []float64, error) {
	var dates []string
	if err := d.db.NewSelect().Model((*ETFDailyPrice)(nil)).Column("trade_date").Where("ts_code=?", code).Order("trade_date DESC").Limit(count).Scan(d.ctx, &dates); err != nil {
		return nil, nil, err
	}

	var closes []float64
	if err := d.db.NewSelect().Model((*ETFDailyPrice)(nil)).Column("close").Where("ts_code=?", code).Order("trade_date DESC").Limit(count).Scan(d.ctx, &closes); err != nil {
		return nil, nil, err
	}
	return lo.Reverse(dates), lo.Reverse(closes), nil
}
