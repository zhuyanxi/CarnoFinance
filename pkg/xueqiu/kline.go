package xueqiu

import (
	"encoding/json"
	"fmt"
	"time"
)

const KLineFormat = "https://stock.xueqiu.com/v5/stock/chart/kline.json?symbol=%s&begin=%d&period=%s&type=%s&count=%d"

type KLine struct {
	Data KLineData `json:"data,omitempty"`
}
type KLineData struct {
	Column           []string    `json:"column,omitempty"`
	Item             [][]float64 `json:"item,omitempty"`
	ErrorCode        int         `json:"error_code,omitempty"`
	ErrorDescription string      `json:"error_description,omitempty"`
}
type KLineQuery struct {
	Symbol string `json:"symbol,omitempty"`
	Period string `json:"period,omitempty"`
	Type   string `json:"type,omitempty"`
	Count  int    `json:"count,omitempty"`
}

// https://stock.xueqiu.com/v5/stock/chart/kline.json?symbol=SZ159915&begin=1709537063110&period=day&type=before&count=-284&indicator=kline,pe,pb,ps,pcf,market_capital,agt,ggt,balance
func (x *XueQiu) GetKline(q KLineQuery) (KLine, error) {
	if x.token == "" {
		return KLine{}, fmt.Errorf("token is empty")
	}

	qUrl := fmt.Sprintf(KLineFormat, q.Symbol, time.Now().UnixMilli(), q.Period, q.Type, q.Count)
	respBody, err := x.requestXueqiu(qUrl)
	if err != nil {
		return KLine{}, err
	}

	var kline KLine
	if err := json.Unmarshal(respBody, &kline); err != nil {
		return KLine{}, err
	}

	return kline, nil
}
