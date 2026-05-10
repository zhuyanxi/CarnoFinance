package db

import (
	"context"
	"database/sql"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"
)

func NewSqlite(ctx context.Context) *bun.DB {
	return openSqlite("finance.db")
}

func NewOptionTrackSqlite(ctx context.Context) *bun.DB {
	return openSqlite(filepath.Join("data", "OptionTrack.db"))
}

func openSqlite(path string) *bun.DB {
	logrus.Infof("opening sqlite3: %s", path)
	if dir := filepath.Dir(path); dir != "." {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			panic(err)
		}
	}
	sqldb, err := sql.Open(sqliteshim.ShimName, path)
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
