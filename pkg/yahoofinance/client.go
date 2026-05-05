package yahoofinance

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	chartURLTemplate  = "https://query1.finance.yahoo.com/v8/finance/chart/%s"
	defaultUserAgent  = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/136.0.0.0 Safari/537.36"
	defaultTimeout    = 20 * time.Second
	defaultMaxRetries = 3
)

type DailyBar struct {
	TradeDate time.Time
	Open      float64
	High      float64
	Low       float64
	Close     float64
}

type Client struct {
	hc         *http.Client
	maxRetries int
	now        func() time.Time
	userAgent  string
}

type chartResponse struct {
	Chart struct {
		Result []chartResult `json:"result"`
		Error  *chartError   `json:"error"`
	} `json:"chart"`
}

type chartResult struct {
	Meta struct {
		ExchangeTimezoneName string `json:"exchangeTimezoneName"`
	} `json:"meta"`
	Timestamp  []int64 `json:"timestamp"`
	Indicators struct {
		Quote []chartQuote `json:"quote"`
	} `json:"indicators"`
}

type chartQuote struct {
	Open  []*float64 `json:"open"`
	High  []*float64 `json:"high"`
	Low   []*float64 `json:"low"`
	Close []*float64 `json:"close"`
}

type chartError struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type httpStatusError struct {
	StatusCode int
	URL        string
	Body       string
}

func (e *httpStatusError) Error() string {
	if e.Body == "" {
		return fmt.Sprintf("yahoo finance request failed: %s returned %d", e.URL, e.StatusCode)
	}
	return fmt.Sprintf("yahoo finance request failed: %s returned %d: %s", e.URL, e.StatusCode, e.Body)
}

func New() *Client {
	return &Client{
		hc:         &http.Client{Timeout: defaultTimeout},
		maxRetries: defaultMaxRetries,
		now:        func() time.Time { return time.Now().UTC() },
		userAgent:  defaultUserAgent,
	}
}

func NormalizeYahooSymbol(symbol string) (string, error) {
	switch {
	case strings.HasPrefix(symbol, "SH") && len(symbol) == 8:
		return symbol[2:] + ".SS", nil
	case strings.HasPrefix(symbol, "SZ") && len(symbol) == 8:
		return symbol[2:] + ".SZ", nil
	case strings.HasPrefix(symbol, "BJ") && len(symbol) == 8:
		return symbol[2:] + ".BJ", nil
	case strings.HasSuffix(symbol, ".SH") && len(symbol) == 9:
		return symbol[:6] + ".SS", nil
	case strings.HasSuffix(symbol, ".SZ") && len(symbol) == 9:
		return symbol[:6] + ".SZ", nil
	case strings.HasSuffix(symbol, ".BJ") && len(symbol) == 9:
		return symbol[:6] + ".BJ", nil
	default:
		return "", fmt.Errorf("unsupported A-share symbol: %s", symbol)
	}
}

func (c *Client) FetchDailyHistory(ctx context.Context, yahooSymbol string) ([]DailyBar, error) {
	var lastErr error
	for attempt := 1; attempt <= c.maxRetries; attempt++ {
		bars, err := c.fetchDailyHistoryOnce(ctx, yahooSymbol)
		if err == nil {
			return bars, nil
		}
		lastErr = err
		if attempt == c.maxRetries || !isRetryable(err) {
			break
		}

		backoff := time.Duration(attempt) * time.Second
		timer := time.NewTimer(backoff)
		select {
		case <-ctx.Done():
			timer.Stop()
			return nil, ctx.Err()
		case <-timer.C:
		}
	}

	return nil, lastErr
}

func (c *Client) fetchDailyHistoryOnce(ctx context.Context, yahooSymbol string) ([]DailyBar, error) {
	reqURL, err := buildChartURL(yahooSymbol, c.now())
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, fmt.Errorf("build yahoo finance request: %w", err)
	}
	req.Header.Set("User-Agent", c.userAgent)
	req.Header.Set("Accept", "application/json")

	resp, err := c.hc.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request yahoo finance: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 512))
		return nil, &httpStatusError{
			StatusCode: resp.StatusCode,
			URL:        reqURL,
			Body:       strings.TrimSpace(string(body)),
		}
	}

	var payload chartResponse
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return nil, fmt.Errorf("decode yahoo finance response for %s: %w", yahooSymbol, err)
	}

	bars, err := dailyBarsFromChartResponse(payload)
	if err != nil {
		return nil, fmt.Errorf("parse yahoo finance response for %s: %w", yahooSymbol, err)
	}

	return bars, nil
}

func dailyBarsFromChartResponse(payload chartResponse) ([]DailyBar, error) {
	if payload.Chart.Error != nil {
		return nil, fmt.Errorf("%s: %s", payload.Chart.Error.Code, payload.Chart.Error.Description)
	}
	if len(payload.Chart.Result) == 0 {
		return nil, errors.New("empty chart result")
	}

	result := payload.Chart.Result[0]
	if len(result.Indicators.Quote) == 0 {
		return nil, errors.New("missing quote data")
	}

	quote := result.Indicators.Quote[0]
	size := minLen(
		len(result.Timestamp),
		len(quote.Open),
		len(quote.High),
		len(quote.Low),
		len(quote.Close),
	)
	if size == 0 {
		return nil, errors.New("no OHLC points returned")
	}

	location := time.UTC
	if result.Meta.ExchangeTimezoneName != "" {
		if loaded, err := time.LoadLocation(result.Meta.ExchangeTimezoneName); err == nil {
			location = loaded
		}
	}

	bars := make([]DailyBar, 0, size)
	for idx := 0; idx < size; idx++ {
		if quote.Open[idx] == nil || quote.High[idx] == nil || quote.Low[idx] == nil || quote.Close[idx] == nil {
			continue
		}

		bars = append(bars, DailyBar{
			TradeDate: time.Unix(result.Timestamp[idx], 0).In(location),
			Open:      *quote.Open[idx],
			High:      *quote.High[idx],
			Low:       *quote.Low[idx],
			Close:     *quote.Close[idx],
		})
	}

	if len(bars) == 0 {
		return nil, errors.New("no complete OHLC points returned")
	}

	return bars, nil
}

func buildChartURL(yahooSymbol string, now time.Time) (string, error) {
	query := make(url.Values)
	query.Set("period1", "0")
	query.Set("period2", fmt.Sprintf("%d", now.Add(24*time.Hour).Unix()))
	query.Set("interval", "1d")
	query.Set("includeAdjustedClose", "false")
	query.Set("events", "div%7Csplit")
	query.Set("lang", "en-US")
	query.Set("region", "US")

	base := fmt.Sprintf(chartURLTemplate, url.PathEscape(yahooSymbol))
	return base + "?" + query.Encode(), nil
}

func isRetryable(err error) bool {
	var statusErr *httpStatusError
	if errors.As(err, &statusErr) {
		return statusErr.StatusCode == http.StatusTooManyRequests || statusErr.StatusCode >= http.StatusInternalServerError
	}
	return false
}

func minLen(values ...int) int {
	if len(values) == 0 {
		return 0
	}

	minValue := values[0]
	for idx := 1; idx < len(values); idx++ {
		if values[idx] < minValue {
			minValue = values[idx]
		}
	}
	return minValue
}
