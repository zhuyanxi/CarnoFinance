package db

import (
	"context"
	"database/sql"

	"github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"
)

func NewSqlite(ctx context.Context) *bun.DB {
	logrus.Infoln("opening sqlite3")
	sqldb, err := sql.Open(sqliteshim.ShimName, "finance.db")
	if err != nil {
		panic(err)
	}
	db := bun.NewDB(sqldb, sqlitedialect.New())
	if err := db.Ping(); err != nil {
		panic(err)
	}
	// db.AddQueryHook(bundebug.NewQueryHook())
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	return db
}
