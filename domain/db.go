package domain

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
)

type Domain struct {
	ctx context.Context
	db  *bun.DB
}

func NewDomain(ctx context.Context, db *bun.DB) *Domain {
	return &Domain{
		ctx: ctx,
		db:  db,
	}
}

func (d *Domain) Init() {
	_, err := d.db.NewCreateTable().Model((*StockDailyPrice)(nil)).Exec(d.ctx)
	logrus.Infof("%+v", err)
	_, err = d.db.NewCreateTable().Model((*ETFDailyPrice)(nil)).Exec(d.ctx)
	logrus.Infof("%+v", err)
}
