package domain

type ETFDailyPrice struct {
	TSCode    string  `json:"ts_code,omitempty" binding:"required"`
	TradeDate string  `json:"trade_date,omitempty"`
	Open      float64 `json:"open,omitempty"`
	High      float64 `json:"high,omitempty"`
	Low       float64 `json:"low,omitempty"`
	Close     float64 `json:"close,omitempty"`
}

func RSRS(prices []ETFDailyPrice, period int) (float64, error) {

	return 0, nil
}

func (d *Domain) InsertETFDailyPrice(data ETFDailyPrice) {
	d.db.NewInsert().Model(&data).Exec(d.ctx)
}
