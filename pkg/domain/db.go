package domain

import (
	"context"
	"errors"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
	"github.com/zhuyanxi/CarnoFinance/pkg/helper"
	"github.com/zhuyanxi/CarnoFinance/pkg/xueqiu"
)

type Domain struct {
	ctx context.Context
	db  *bun.DB
	xqc *xueqiu.XueQiu
}

func NewDomain(ctx context.Context, db *bun.DB, xqc *xueqiu.XueQiu) *Domain {
	return &Domain{
		ctx: ctx,
		db:  db,
		xqc: xqc,
	}
}

func (d *Domain) Init() {
	_, err := d.db.NewCreateTable().Model((*StockDailyPrice)(nil)).Exec(d.ctx)
	logrus.Infof("%+v", err)
	_, err = d.db.NewCreateTable().Model((*ETFDailyPrice)(nil)).Exec(d.ctx)
	logrus.Infof("%+v", err)
	_, err = d.db.NewCreateTable().Model((*ETFCodeList)(nil)).Exec(d.ctx)
	logrus.Infof("%+v", err)
}

// {
//     "data": {
//         "symbol": "SZ159915",
//         "column": [
//             "timestamp",
//             "volume",
//             "open",
//             "high",
//             "low",
//             "close"
//         ],
//         "item": [
//             [
//                 1709222400000,
//                 1288194720,
//                 1.757,
//                 1.781,
//                 1.746,
//                 1.775
//             ]
//         ]
//     },
//     "error_code": 0,
//     "error_description": ""
// }

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

	item := kline.Data.Item[0]
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

	return nil
}
