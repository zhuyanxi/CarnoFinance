package db

import (
	"database/sql"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strconv"

	_ "github.com/duckdb/duckdb-go/v2"
)

func NewDuckDB(dbPath string, threads int) (*sql.DB, error) {
	if dbPath != "" {
		parentDir := filepath.Dir(dbPath)
		if parentDir != "." {
			if err := os.MkdirAll(parentDir, 0o755); err != nil {
				return nil, fmt.Errorf("create duckdb parent dir: %w", err)
			}
		}
	}

	dsn := buildDuckDBDSN(dbPath, threads)
	sqldb, err := sql.Open("duckdb", dsn)
	if err != nil {
		return nil, fmt.Errorf("open duckdb: %w", err)
	}

	sqldb.SetMaxOpenConns(1)
	sqldb.SetMaxIdleConns(1)

	if err := sqldb.Ping(); err != nil {
		sqldb.Close()
		return nil, fmt.Errorf("ping duckdb: %w", err)
	}

	return sqldb, nil
}

func buildDuckDBDSN(dbPath string, threads int) string {
	query := make(url.Values)
	query.Set("access_mode", "READ_WRITE")
	if threads > 0 {
		query.Set("threads", strconv.Itoa(threads))
	}

	if dbPath == "" {
		return "?" + query.Encode()
	}

	return dbPath + "?" + query.Encode()
}
