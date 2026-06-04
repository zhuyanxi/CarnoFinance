package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

const (
	eastmoneyKlineURL = "https://push2his.eastmoney.com/api/qt/stock/kline/get?secid=1.000688&fields1=f1,f2,f3,f4,f5,f6&fields2=f51,f52,f53,f54,f55&klt=101&fqt=0&beg=20191231&end=20500101"
	outputFileName    = "kc50_daily.csv"
	requestTimeout    = 30 * time.Second
	maxRetries        = 3
)

type eastmoneyKlineResponse struct {
	Data *struct {
		Code   string   `json:"code"`
		Name   string   `json:"name"`
		Klines []string `json:"klines"`
	} `json:"data"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	outputPath, err := resolveOutputPath()
	if err != nil {
		fmt.Fprintf(os.Stderr, "resolve output path: %v\n", err)
		os.Exit(1)
	}

	rows, err := fetchKC50Rows(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch KC50 history: %v\n", err)
		os.Exit(1)
	}

	if err := os.MkdirAll(filepath.Dir(outputPath), 0o755); err != nil {
		fmt.Fprintf(os.Stderr, "create output dir: %v\n", err)
		os.Exit(1)
	}

	file, err := os.Create(outputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "create %s: %v\n", outputPath, err)
		os.Exit(1)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.Write([]string{"trade_date", "open", "high", "low", "close"}); err != nil {
		fmt.Fprintf(os.Stderr, "write csv header: %v\n", err)
		os.Exit(1)
	}

	for _, row := range rows {
		if err := writer.Write(row); err != nil {
			fmt.Fprintf(os.Stderr, "write csv row: %v\n", err)
			os.Exit(1)
		}
	}

	if err := writer.Error(); err != nil {
		fmt.Fprintf(os.Stderr, "flush csv writer: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("saved %d rows to %s\n", len(rows), outputPath)
}

func resolveOutputPath() (string, error) {
	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		return "", fmt.Errorf("cannot resolve current file path")
	}
	return filepath.Join(filepath.Dir(currentFile), outputFileName), nil
}

func fetchKC50Rows(ctx context.Context) ([][]string, error) {
	var payload eastmoneyKlineResponse
	var lastErr error
	for attempt := 1; attempt <= maxRetries; attempt++ {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, eastmoneyKlineURL, nil)
		if err != nil {
			return nil, err
		}
		req.Header.Set("User-Agent", "Mozilla/5.0")
		req.Header.Set("Referer", "https://quote.eastmoney.com")

		resp, err := (&http.Client{Timeout: requestTimeout}).Do(req)
		if err != nil {
			lastErr = err
		} else {
			func() {
				defer resp.Body.Close()
				if resp.StatusCode != http.StatusOK {
					lastErr = fmt.Errorf("unexpected status: %s", resp.Status)
					return
				}
				if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
					lastErr = err
					return
				}
				lastErr = nil
			}()
		}
		if lastErr == nil {
			break
		}
		if attempt < maxRetries {
			timer := time.NewTimer(time.Duration(attempt) * time.Second)
			select {
			case <-ctx.Done():
				timer.Stop()
				return nil, ctx.Err()
			case <-timer.C:
			}
		}
	}
	if lastErr != nil {
		return nil, lastErr
	}

	if payload.Data == nil || len(payload.Data.Klines) == 0 {
		return nil, fmt.Errorf("empty Eastmoney kline payload")
	}

	rows := make([][]string, 0, len(payload.Data.Klines))
	for _, line := range payload.Data.Klines {
		parts := strings.Split(line, ",")
		if len(parts) < 5 {
			return nil, fmt.Errorf("unexpected kline format: %s", line)
		}
		rows = append(rows, []string{parts[0], parts[1], parts[3], parts[4], parts[2]})
	}

	return rows, nil
}