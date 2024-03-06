package domain

import (
	"sort"

	"github.com/samber/lo"
	"github.com/sirupsen/logrus"
	"github.com/zhuyanxi/CarnoFinance/pkg/helper"
)

type ETFCodeList struct {
	TSCode string `json:"ts_code,omitempty"`
	Name   string `json:"name,omitempty"`
}

type ETFDailyPrice struct {
	TSCode    string  `json:"ts_code,omitempty" binding:"required"`
	TradeDate string  `json:"trade_date,omitempty"`
	Open      float64 `json:"open,omitempty"`
	High      float64 `json:"high,omitempty"`
	Low       float64 `json:"low,omitempty"`
	Close     float64 `json:"close,omitempty"`
}

type RSRSDto struct {
	TsCode string  `json:"ts_code,omitempty"`
	Name   string  `json:"name,omitempty"`
	RSRS   float64 `json:"rsrs,omitempty"`
}

func (d *Domain) GetETFCodeList() ([]ETFCodeList, error) {
	var codes []ETFCodeList
	err := d.db.NewSelect().Model(&codes).Scan(d.ctx)
	if err != nil {
		return nil, err
	}
	return codes, nil
}

func (d *Domain) GetRSRSList(period int, order string) ([]RSRSDto, error) {
	etfList, err := d.GetETFCodeList()
	if err != nil {
		return nil, err
	}

	var rets []RSRSDto
	for _, etf := range etfList {
		var pricesList []float64
		err := d.db.NewSelect().Model((*ETFDailyPrice)(nil)).Column("close").
			Where("ts_code=?", etf.TSCode).
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

func (d *Domain) InsertETFDailyPrice(data ETFDailyPrice) {
	exist, _ := d.db.NewSelect().Model((*ETFDailyPrice)(nil)).Where("ts_code=?", data.TSCode).Where("trade_date=?", data.TradeDate).Exists(d.ctx)
	if exist {
		logrus.Infof("ETFDailyPrice %s %s already exists", data.TSCode, data.TradeDate)
		return
	}
	d.db.NewInsert().Model(&data).Exec(d.ctx)
}
