package tushare

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type StockBasicRequest struct {
	TSCode     string `json:"ts_code,omitempty"`
	Exchange   string `json:"exchange,omitempty"`
	Market     string `json:"market,omitempty"`
	ListStatus string `json:"list_status,omitempty"`
	IsHS       string `json:"is_hs,omitempty"`
}

type StockBasicItem struct {
	TSCode     string
	Symbol     string
	Name       string
	Area       string
	Industry   string
	Market     string
	ListStatus string
	ListDate   string
}

func (t *Tushare) GetStockBasic(stockReq StockBasicRequest) ([]StockBasicItem, error) {
	tushareReq := TushareRequest{
		APIName: "stock_basic",
		Token:   t.Token,
		Params:  stockReq,
		Fields:  "ts_code,symbol,name,area,industry,market,list_status,list_date",
	}

	payload, err := json.Marshal(tushareReq)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, TuShareDomain, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	tushareResp := new(TushareResponse)
	if err := json.Unmarshal(body, tushareResp); err != nil {
		return nil, err
	}
	if tushareResp.Code != 0 {
		return nil, fmt.Errorf("get stock basic failed: %s", tushareResp.Msg)
	}

	return convertToStockBasicItems(tushareResp.Data), nil
}

func convertToStockBasicItems(data TushareResponseData) []StockBasicItem {
	indexByField := make(map[string]int, len(data.Fields))
	for idx, field := range data.Fields {
		indexByField[field] = idx
	}

	items := make([]StockBasicItem, 0, len(data.Items))
	for _, row := range data.Items {
		items = append(items, StockBasicItem{
			TSCode:     readStringField(row, indexByField, "ts_code"),
			Symbol:     readStringField(row, indexByField, "symbol"),
			Name:       readStringField(row, indexByField, "name"),
			Area:       readStringField(row, indexByField, "area"),
			Industry:   readStringField(row, indexByField, "industry"),
			Market:     readStringField(row, indexByField, "market"),
			ListStatus: readStringField(row, indexByField, "list_status"),
			ListDate:   readStringField(row, indexByField, "list_date"),
		})
	}

	return items
}

func readStringField(row []any, indexByField map[string]int, field string) string {
	idx, ok := indexByField[field]
	if !ok || idx >= len(row) {
		return ""
	}

	value, ok := row[idx].(string)
	if !ok {
		return ""
	}
	return value
}
