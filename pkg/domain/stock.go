package domain

import "github.com/sirupsen/logrus"

type StockCodeList struct {
	TSCode string `bun:"ts_code,pk" json:"ts_code,omitempty"`
	Name   string `json:"name,omitempty"`
}

type StockDailyPrice struct {
	TSCode    string  `json:"ts_code,omitempty"`
	TradeDate string  `json:"trade_date,omitempty"`
	Open      float64 `json:"open,omitempty"`
	High      float64 `json:"high,omitempty"`
	Low       float64 `json:"low,omitempty"`
	Close     float64 `json:"close,omitempty"`
	PreClose  float64 `json:"pre_close,omitempty"`
	Change    float64 `json:"change,omitempty"`
	PctChg    float64 `json:"pct_chg,omitempty"`
	Vol       float64 `json:"vol,omitempty"`
	Amount    float64 `json:"amount,omitempty"`
}

func (d *Domain) InitStockList() {
	stocks, err := d.xqc.GetStockList(1, 10000)
	if err != nil {
		logrus.Errorln("init stock list failed: ", err.Error())
		return
	}

	var codeList []StockCodeList
	for _, stock := range stocks.Data.List {
		codeList = append(codeList, StockCodeList{
			TSCode: stock.Symbol,
			Name:   stock.Name,
		})
	}

	var num int64
	if len(codeList) != 0 {
		res, _ := d.db.NewInsert().Model(&codeList).Ignore().Exec(d.ctx)
		num, _ = res.RowsAffected()
	}
	logrus.Infof("init stock list row nums: %d", num)
}
