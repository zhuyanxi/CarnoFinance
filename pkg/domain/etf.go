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
		d.SetETFPrice(code.TSCode, -5)
	}
	return nil
}

func (d *Domain) SetETFPrice(code string, recentCount int) error {
	kline, err := d.xqc.GetKline(xueqiu.KLineQuery{
		Symbol: code,
		Period: "day",
		Type:   xueqiu.KLineTypeBefore,
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

func (d *Domain) CountHighGreaterThanOpen(tsCode string) (count int, proportion float64, err error) {
	total, err := d.db.NewSelect().Model(new(ETFDailyPrice)).Where("ts_code = ?", tsCode).Count(d.ctx)
	if err != nil {
		logrus.Errorf("failed to count total records for ts_code %s: %v", tsCode, err)
		return 0, 0, err
	}

	if total == 0 {
		logrus.Warnf("no records found for ts_code %s", tsCode)
		return 0, 0, nil
	}

	count, err = d.db.NewSelect().Model(new(ETFDailyPrice)).Where("ts_code = ?", tsCode).Where("high > open").Count(d.ctx)
	if err != nil {
		logrus.Errorf("failed to count high > open records for ts_code %s: %v", tsCode, err)
		return 0, 0, err
	}

	proportion = float64(count) / float64(total)
	return count, proportion, nil
}

// gap=0.05,total=1.728
// gap=0.051,total=1.786
// gap=0.052,total=1.854
// Algo:
// 黄金ETFSH518880交易: 每天开盘以开盘价买入, 如果最高价高于开盘价+0.02, 则在当天以开盘价+0.02卖出, 否则在当天以收盘价卖出
func (d *Domain) CalcGoldETFTotalProfit() (profit float64) {
	gap := 0.1
	dayEarn := 0
	dayLoss := 0
	daySame := 0
	records := make([]ETFDailyPrice, 0)
	err := d.db.NewSelect().Model(&records).Where("ts_code = ?", "SH518880").Order("trade_date ASC").Scan(d.ctx)
	if err != nil {
		logrus.Errorf("failed to fetch records for SH518880: %v", err)
		return 0
	}

	for _, record := range records {
		buyPrice := record.Open
		sellPrice := record.Close

		if record.High > buyPrice+gap {
			sellPrice = buyPrice + gap
		}

		if sellPrice > buyPrice {
			dayEarn++
		} else if sellPrice < buyPrice {
			dayLoss++
		} else {
			daySame++
		}

		profit += sellPrice - buyPrice
	}

	logrus.Infof("Total profit: %f, Earn days: %d, Loss days: %d, Same days: %d", profit, dayEarn, dayLoss, daySame)
	return profit
}
