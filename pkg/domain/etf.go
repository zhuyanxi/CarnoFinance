package domain

import (
	"errors"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/zhuyanxi/CarnoFinance/pkg/helper"
	"github.com/zhuyanxi/CarnoFinance/pkg/xueqiu"
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

func (d *Domain) GetETFCodeList() ([]ETFCodeList, error) {
	var codes []ETFCodeList
	err := d.db.NewSelect().Model(&codes).Scan(d.ctx)
	if err != nil {
		return nil, err
	}
	return codes, nil
}

func (d *Domain) InsertETFDailyPrice(data ETFDailyPrice) {
	exist, _ := d.db.NewSelect().Model((*ETFDailyPrice)(nil)).Where("ts_code=?", data.TSCode).Where("trade_date=?", data.TradeDate).Exists(d.ctx)
	if exist {
		logrus.Infof("ETFDailyPrice %s %s already exists", data.TSCode, data.TradeDate)
		return
	}
	d.db.NewInsert().Model(&data).Exec(d.ctx)
}

func (d *Domain) InitLastOneDayETFPrice() error {
	codes, err := d.GetETFCodeList()
	if err != nil {
		logrus.Errorf("get etf code list error: %v", err)
		return err
	}
	for _, code := range codes {
		d.SetETFPrice(code.TSCode, -1)
	}
	return nil
}

func (d *Domain) SetETFPrice(code string, recentCount int) error {
	kline, err := d.xqc.GetKline(xueqiu.KLineQuery{
		Symbol: code,
		Period: "day",
		Type:   "before",
		Count:  recentCount,
	})
	if err != nil {
		logrus.Errorf("set etf last price failed: %+v", err)
		return err
	}

	if len(kline.Data.Item) == 0 {
		logrus.Warnln("kline is empty")
		return errors.New("kline is empty")
	}

	for i := 0; i < len(kline.Data.Item); i++ {
		item := kline.Data.Item[i]
		ts := int64(item[0])
		open := item[2]
		high := item[3]
		low := item[4]
		close := item[5]

		var etf ETFDailyPrice
		etf.TSCode = code
		etf.Open = open
		etf.High = high
		etf.Low = low
		etf.Close = close
		etf.TradeDate = time.Unix(ts/1000, ts%1000).Format(helper.DateFormat)
		d.InsertETFDailyPrice(etf)
	}

	return nil
}
