package main

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sort"
	"strings"
	"sync"
	"syscall"
	"time"

	duckdb "github.com/duckdb/duckdb-go/v2"
	"github.com/sirupsen/logrus"
	dbpkg "github.com/zhuyanxi/CarnoFinance/pkg/db"
	"github.com/zhuyanxi/CarnoFinance/pkg/sina"
	"github.com/zhuyanxi/CarnoFinance/pkg/yahoofinance"
)

const (
	defaultDBPath      = "data/all-stock.db"
	defaultConcurrency = 16
	defaultDuckThreads = 4
)

var dailyPriceAppenderColumns = []string{"symbol", "yahoo_symbol", "trade_date", "open", "high", "low", "close"}

type company struct {
	Symbol      string
	Name        string
	MarketType  string
	YahooSymbol string
}

type fetchResult struct {
	Company company
	Bars    []yahoofinance.DailyBar
	Err     error
}

type importSummary struct {
	Success int
	Failed  int
	Rows    int
}

type duckStore struct {
	db *sql.DB
}

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})

	dbPath := flag.String("db", defaultDBPath, "DuckDB database path")
	concurrency := flag.Int("concurrency", defaultConcurrency, "Number of concurrent Yahoo Finance workers")
	limit := flag.Int("limit", 0, "Limit the number of symbols to import, 0 means all")
	symbols := flag.String("symbols", "", "Comma-separated A-share ts_code values to import, e.g. 600519.SH,000001.SZ")
	duckThreads := flag.Int("duckdb-threads", defaultDuckThreads, "DuckDB execution threads")
	flag.Parse()
	requestedSymbols := splitCSV(*symbols)

	if *concurrency <= 0 {
		logrus.Fatalln("concurrency must be positive")
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	var companies []company
	var err error
	if len(requestedSymbols) != 0 {
		companies, err = buildCompaniesFromSymbols(requestedSymbols)
		if err != nil {
			logrus.Fatalf("build import list: %v", err)
		}
	} else {
		sinaClient := sina.New()
		companies, err = loadCompanies(ctx, sinaClient)
		if err != nil {
			logrus.Fatalf("load A-share company list: %v", err)
		}
	}

	companies, err = filterCompanies(companies, requestedSymbols, *limit)
	if err != nil {
		logrus.Fatalf("prepare import list: %v", err)
	}
	if len(companies) == 0 {
		logrus.Infoln("no companies matched the requested filters")
		return
	}

	store, err := newDuckStore(*dbPath, *duckThreads)
	if err != nil {
		logrus.Fatalf("open duckdb: %v", err)
	}
	defer store.Close()

	if err := store.Init(ctx); err != nil {
		logrus.Fatalf("initialize duckdb schema: %v", err)
	}

	client := yahoofinance.New()
	results := make(chan fetchResult, *concurrency)
	jobs := make(chan company)

	var workers sync.WaitGroup
	for idx := 0; idx < *concurrency; idx++ {
		workers.Add(1)
		go func() {
			defer workers.Done()
			fetchWorker(ctx, client, jobs, results)
		}()
	}

	go func() {
		defer close(jobs)
		for _, item := range companies {
			select {
			case <-ctx.Done():
				return
			case jobs <- item:
			}
		}
	}()

	go func() {
		workers.Wait()
		close(results)
	}()

	summary := importSummary{}
	total := len(companies)
	processed := 0
	for result := range results {
		processed++
		if result.Err != nil {
			summary.Failed++
			logrus.Errorf("[%d/%d] %s (%s) failed: %v", processed, total, result.Company.Symbol, result.Company.YahooSymbol, result.Err)
			continue
		}

		if err := store.UpsertHistory(ctx, result.Company, result.Bars); err != nil {
			summary.Failed++
			logrus.Errorf("[%d/%d] persist %s failed: %v", processed, total, result.Company.Symbol, err)
			continue
		}

		summary.Success++
		summary.Rows += len(result.Bars)
		if processed%25 == 0 || processed == total {
			logrus.Infof("progress: %d/%d symbols imported, success=%d failed=%d rows=%d", processed, total, summary.Success, summary.Failed, summary.Rows)
		}
	}

	if err := ctx.Err(); err != nil && err != context.Canceled {
		logrus.Fatalf("import interrupted: %v", err)
	}

	logrus.Infof("import finished: symbols=%d success=%d failed=%d rows=%d db=%s", total, summary.Success, summary.Failed, summary.Rows, *dbPath)
}

func loadCompanies(ctx context.Context, sinaClient *sina.Client) ([]company, error) {
	items, err := sinaClient.GetAShareList(ctx)
	if err != nil {
		return nil, err
	}

	companies := make([]company, 0, len(items))
	skippedUnsupported := 0
	for _, item := range items {
		if item.Exchange == "BJ" {
			skippedUnsupported++
			continue
		}

		tsCode := item.Code + "." + item.Exchange
		yahooSymbol, err := yahoofinance.NormalizeYahooSymbol(tsCode)
		if err != nil {
			logrus.Warnf("skip unsupported ts_code %s: %v", tsCode, err)
			continue
		}

		companies = append(companies, company{
			Symbol:      tsCode,
			Name:        item.Name,
			MarketType:  item.MarketType,
			YahooSymbol: yahooSymbol,
		})
	}
	sort.Slice(companies, func(left, right int) bool {
		return companies[left].Symbol < companies[right].Symbol
	})
	if skippedUnsupported > 0 {
		logrus.Infof("skipped Yahoo-unsupported BJ symbols: count=%d", skippedUnsupported)
	}
	logrus.Infof("loaded Yahoo-supported listed A-share companies from Sina: count=%d", len(companies))
	return companies, nil
}

func buildCompaniesFromSymbols(symbols []string) ([]company, error) {
	companies := make([]company, 0, len(symbols))
	for _, symbol := range symbols {
		yahooSymbol, err := yahoofinance.NormalizeYahooSymbol(symbol)
		if err != nil {
			return nil, err
		}
		companies = append(companies, company{
			Symbol:      symbol,
			Name:        symbol,
			MarketType:  inferMarketType(symbol),
			YahooSymbol: yahooSymbol,
		})
	}
	return companies, nil
}

func filterCompanies(companies []company, symbols []string, limit int) ([]company, error) {
	if len(symbols) != 0 {
		allowed := make(map[string]struct{}, len(symbols))
		for _, symbol := range symbols {
			allowed[symbol] = struct{}{}
		}

		filtered := make([]company, 0, len(symbols))
		for _, item := range companies {
			if _, ok := allowed[item.Symbol]; ok {
				filtered = append(filtered, item)
			}
		}
		companies = filtered
	}

	if limit > 0 && len(companies) > limit {
		companies = companies[:limit]
	}

	return companies, nil
}

func fetchWorker(ctx context.Context, client *yahoofinance.Client, jobs <-chan company, results chan<- fetchResult) {
	for {
		select {
		case <-ctx.Done():
			return
		case item, ok := <-jobs:
			if !ok {
				return
			}

			bars, err := client.FetchDailyHistory(ctx, item.YahooSymbol)
			result := fetchResult{Company: item, Bars: bars, Err: err}

			select {
			case <-ctx.Done():
				return
			case results <- result:
			}
		}
	}
}

func newDuckStore(dbPath string, threads int) (*duckStore, error) {
	sqldb, err := dbpkg.NewDuckDB(dbPath, threads)
	if err != nil {
		return nil, err
	}
	return &duckStore{db: sqldb}, nil
}

func (s *duckStore) Init(ctx context.Context) error {
	statements := []string{
		`CREATE TABLE IF NOT EXISTS a_share_companies (
			symbol VARCHAR PRIMARY KEY,
			yahoo_symbol VARCHAR NOT NULL,
			name VARCHAR NOT NULL,
			market_type VARCHAR NOT NULL,
			updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS a_share_daily_prices (
			symbol VARCHAR NOT NULL,
			yahoo_symbol VARCHAR NOT NULL,
			trade_date DATE NOT NULL,
			open DOUBLE NOT NULL,
			high DOUBLE NOT NULL,
			low DOUBLE NOT NULL,
			close DOUBLE NOT NULL,
			updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY(symbol, trade_date)
		)`,
		`CREATE INDEX IF NOT EXISTS idx_a_share_daily_prices_trade_date ON a_share_daily_prices(trade_date)`,
	}

	for _, stmt := range statements {
		if _, err := s.db.ExecContext(ctx, stmt); err != nil {
			return err
		}
	}
	return nil
}

func (s *duckStore) UpsertHistory(ctx context.Context, item company, bars []yahoofinance.DailyBar) error {
	conn, err := s.db.Conn(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	if _, err := conn.ExecContext(ctx, `
		INSERT INTO a_share_companies (
			symbol, yahoo_symbol, name, market_type
		) VALUES (?, ?, ?, ?)
		ON CONFLICT(symbol) DO UPDATE SET
			yahoo_symbol = excluded.yahoo_symbol,
			name = excluded.name,
			market_type = excluded.market_type,
			updated_at = now()
	`, item.Symbol, item.YahooSymbol, item.Name, item.MarketType); err != nil {
		return err
	}

	if _, err := conn.ExecContext(ctx, `DELETE FROM a_share_daily_prices WHERE symbol = ?`, item.Symbol); err != nil {
		return err
	}

	return conn.Raw(func(driverConn any) error {
		rawConn, ok := driverConn.(driver.Conn)
		if !ok {
			return fmt.Errorf("unexpected duckdb raw conn type %T", driverConn)
		}

		appender, err := duckdb.NewAppenderWithColumns(rawConn, "", "", "a_share_daily_prices", dailyPriceAppenderColumns)
		if err != nil {
			return err
		}

		for _, bar := range bars {
			tradeDate := time.Date(bar.TradeDate.Year(), bar.TradeDate.Month(), bar.TradeDate.Day(), 0, 0, 0, 0, time.UTC)
			if err := appender.AppendRow(item.Symbol, item.YahooSymbol, tradeDate, bar.Open, bar.High, bar.Low, bar.Close); err != nil {
				_ = appender.Clear()
				_ = appender.Close()
				return err
			}
		}

		return appender.Close()
	})
}

func (s *duckStore) Close() error {
	return s.db.Close()
}

func splitCSV(value string) []string {
	parts := strings.Split(value, ",")
	items := make([]string, 0, len(parts))
	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed == "" {
			continue
		}
		items = append(items, trimmed)
	}
	return items
}

func inferMarketType(symbol string) string {
	switch {
	case strings.HasSuffix(symbol, ".SH") || strings.HasPrefix(symbol, "SH"):
		return "SH"
	case strings.HasSuffix(symbol, ".SZ") || strings.HasPrefix(symbol, "SZ"):
		return "SZ"
	case strings.HasSuffix(symbol, ".BJ") || strings.HasPrefix(symbol, "BJ"):
		return "BJ"
	default:
		return ""
	}
}
