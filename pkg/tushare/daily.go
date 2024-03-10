package tushare

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/zhuyanxi/CarnoFinance/pkg/domain"
	"github.com/zhuyanxi/CarnoFinance/pkg/utils"
)

type DailyRequest struct {
	TSCode    string `json:"ts_code,omitempty"`
	TradeDate string `json:"trade_date,omitempty"`
	StartDate string `json:"start_date,omitempty"`
	EndDate   string `json:"end_date,omitempty"`
}

// type DailyResponse struct {
// 	TSCode    string  `json:"ts_code,omitempty"`
// 	TradeDate string  `json:"trade_date,omitempty"`
// 	Open      float64 `json:"open,omitempty"`
// 	High      float64 `json:"high,omitempty"`
// 	Low       float64 `json:"low,omitempty"`
// 	Close     float64 `json:"close,omitempty"`
// 	PreClose  float64 `json:"pre_close,omitempty"`
// 	Change    float64 `json:"change,omitempty"`
// 	PctChg    float64 `json:"pct_chg,omitempty"`
// 	Vol       float64 `json:"vol,omitempty"`
// 	Amount    float64 `json:"amount,omitempty"`
// }

func (t *Tushare) GetDailyPrices(dailyReq DailyRequest) ([]domain.StockDailyBeforePrice, error) {
	tushareReq := TushareRequest{
		APIName: "daily",
		Token:   t.Token,
		Params:  dailyReq,
	}

	payload, err := json.Marshal(tushareReq)
	if err != nil {
		logrus.Errorln("Error marshal request:", err)
		return nil, err
	}

	// payload := []byte(`{
	// 	"api_name": "daily",
	// 	"token": "aa",
	// 	"params": {
	// 		"ts_code": "000001.SZ",
	// 		"start_date": "20231101",
	// 		"end_date": "20231118",
	// 		"is_open": "0"
	// 	},
	// 	"fields": "exchange,cal_date,is_open,pretrade_date"
	// }`)

	req, err := http.NewRequest(http.MethodPost, TuShareDomain, bytes.NewBuffer(payload))
	if err != nil {
		logrus.Errorln("Error creating request:", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logrus.Errorln("Error sending request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Errorln("Error reading response body:", err)
		return nil, err
	}
	//"{\"code\":40101,\"msg\":\"抱歉，您输入的TOKEN无效！\",\"message\":\"抱歉，您输入的TOKEN无效！\",\"data\":null}"

	//"{\"request_id\":\"48b42110868811ee970e05c3e65b8164\",\"code\":0,\"msg\":\"\",\"data\":{\"fields\":[\"ts_code\",\"trade_date\",\"open\",\"high\",\"low\",\"close\",\"pre_close\",\"change\",\"pct_chg\",\"vol\",\"amount\"],\"items\":[[\"000001.SZ\",\"20231117\",10.22,10.25,10.12,10.15,10.24,-0.09,-0.8789,903362.17,917616.681],[\"000001.SZ\",\"20231116\",10.38,10.39,10.23,10.24,10.37,-0.13,-1.2536,922621.94,948395.097],[\"000001.SZ\",\"20231115\",10.37,10.45,10.34,10.37,10.27,0.1,0.9737,1109140.15,1152154.78],[\"000001.SZ\",\"20231114\",10.24,10.33,10.23,10.27,10.2"
	tushareResp := new(TushareResponse)
	if err := json.Unmarshal(body, &tushareResp); err != nil {
		logrus.Errorln("Error unmarshal response body:", err)
		return nil, err
	}
	if tushareResp.Code != 0 {
		return nil, fmt.Errorf("get daily prices failed: %s", tushareResp.Msg)
	}

	return convertToDailyPrices(tushareResp.Data), nil
}

func convertToDailyPrices(data TushareResponseData) []domain.StockDailyBeforePrice {
	rets := make([]domain.StockDailyBeforePrice, 0)

	for _, item := range data.Items {
		dr := new(domain.StockDailyBeforePrice)
		utils.SetStructFieldValuesByJsonTags(dr, data.Fields, item)
		rets = append(rets, *dr)
	}

	return rets
}
