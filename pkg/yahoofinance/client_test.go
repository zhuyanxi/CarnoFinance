package yahoofinance

import (
	"testing"
	"time"
)

func TestNormalizeYahooSymbol(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		symbol  string
		want    string
		wantErr bool
	}{
		{name: "Shanghai", symbol: "SH600519", want: "600519.SS"},
		{name: "Shenzhen", symbol: "SZ000001", want: "000001.SZ"},
		{name: "Beijing", symbol: "BJ430047", want: "430047.BJ"},
		{name: "Shanghai TSCode", symbol: "600519.SH", want: "600519.SS"},
		{name: "Shenzhen TSCode", symbol: "000001.SZ", want: "000001.SZ"},
		{name: "Beijing TSCode", symbol: "430047.BJ", want: "430047.BJ"},
		{name: "Invalid", symbol: "HK00700", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := NormalizeYahooSymbol(tt.symbol)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("NormalizeYahooSymbol(%q) expected error", tt.symbol)
				}
				return
			}

			if err != nil {
				t.Fatalf("NormalizeYahooSymbol(%q) unexpected error: %v", tt.symbol, err)
			}
			if got != tt.want {
				t.Fatalf("NormalizeYahooSymbol(%q) = %q, want %q", tt.symbol, got, tt.want)
			}
		})
	}
}

func TestDailyBarsFromChartResponse(t *testing.T) {
	t.Parallel()

	open1 := 10.1
	high1 := 10.8
	low1 := 9.9
	close1 := 10.5
	open2 := 10.6
	high2 := 11.0
	low2 := 10.4
	close2 := 10.9

	payload := chartResponse{}
	payload.Chart.Result = []chartResult{{
		Timestamp: []int64{1717353600, 1717440000},
		Indicators: struct {
			Quote []chartQuote `json:"quote"`
		}{
			Quote: []chartQuote{{
				Open:  []*float64{&open1, &open2},
				High:  []*float64{&high1, &high2},
				Low:   []*float64{&low1, &low2},
				Close: []*float64{&close1, &close2},
			}},
		},
	}}
	payload.Chart.Result[0].Meta.ExchangeTimezoneName = "Asia/Shanghai"

	bars, err := dailyBarsFromChartResponse(payload)
	if err != nil {
		t.Fatalf("dailyBarsFromChartResponse() unexpected error: %v", err)
	}
	if len(bars) != 2 {
		t.Fatalf("dailyBarsFromChartResponse() returned %d bars, want 2", len(bars))
	}

	if got := bars[0].TradeDate.Format("2006-01-02"); got != "2024-06-03" {
		t.Fatalf("first trade date = %s, want 2024-06-03", got)
	}
	if bars[1].Close != close2 {
		t.Fatalf("second close = %f, want %f", bars[1].Close, close2)
	}
	if bars[0].TradeDate.Location().String() != time.FixedZone("", 0).String() && bars[0].TradeDate.Location().String() != "Asia/Shanghai" {
		t.Fatalf("unexpected trade date location %s", bars[0].TradeDate.Location())
	}
}
