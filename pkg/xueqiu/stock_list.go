package xueqiu

import (
	"encoding/json"
	"fmt"
)

const StockListFormat = "https://stock.xueqiu.com/v5/stock/screener/quote/list.json?page=%d&size=%d&order=desc&order_by=market_capital&market=CN&type=sh_sz"

type StockList struct {
	Data             StockListData `json:"data,omitempty"`
	ErrorCode        int           `json:"error_code,omitempty"`
	ErrorDescription string        `json:"error_description,omitempty"`
}
type StockListData struct {
	Count int             `json:"count,omitempty"`
	List  []StockListItem `json:"list,omitempty"`
}
type StockListItem struct {
	Symbol string `json:"symbol,omitempty"`
	Name   string `json:"name,omitempty"`
}

func (x *XueQiu) GetStockList(page, size int) (StockList, error) {
	// if x.token == "" {
	// 	return StockList{}, fmt.Errorf("token is empty")
	// }

	qUrl := fmt.Sprintf(StockListFormat, page, size)
	respBody, err := x.requestXueqiu(qUrl)
	if err != nil {
		return StockList{}, err
	}

	var stockList StockList
	if err := json.Unmarshal(respBody, &stockList); err != nil {
		return StockList{}, err
	}
	return stockList, nil
}
