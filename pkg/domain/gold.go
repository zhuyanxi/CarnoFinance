package domain

type GoldETFHistory struct {
	Date  string  `json:"date,omitempty"`
	Open  float64 `json:"open,omitempty"`
	High  float64 `json:"high,omitempty"`
	Low   float64 `json:"low,omitempty"`
	Close float64 `json:"close,omitempty"`
}
