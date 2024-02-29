package domain

type StockDailyPrice struct {
	TSCode    string  `json:"ts_code,omitempty"`
	TradeDate string  `json:"trade_date,omitempty"`
	Open      float64 `json:"open,omitempty"`
	High      float64 `json:"high,omitempty"`
	Low       float64 `json:"low,omitempty"`
	Close     float64 `json:"close,omitempty"`
	PreClose  float64 `json:"pre_close,omitempty"`
	Change    float64 `json:"change,omitempty"`
	PctChg    float64 `json:"pct_chg,omitempty"`
	Vol       float64 `json:"vol,omitempty"`
	Amount    float64 `json:"amount,omitempty"`
}
