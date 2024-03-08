package domain

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
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
	_, err = d.db.NewCreateTable().Model((*StockCodeList)(nil)).Exec(d.ctx)
	logrus.Infof("%+v", err)
	_, err = d.db.NewCreateTable().Model((*ETFCodeList)(nil)).Exec(d.ctx)
	logrus.Infof("%+v", err)

	d.InitLastOneDayETFPrice()
	d.InitStockList()
}
