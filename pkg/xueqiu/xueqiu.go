package xueqiu

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	XueQiuHost  = "https://xueqiu.com/"
	UserAgent   = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/000000000 Safari/537.36"
	TokenName   = "xq_a_token"
	KLineFormat = "https://stock.xueqiu.com/v5/stock/chart/kline.json?symbol=%s&begin=%d&period=%s&type=%s&count=%d"
)

type XueQiu struct {
	hc *http.Client

	token string
}

func New() *XueQiu {
	xq := new(XueQiu)
	client := &http.Client{}

	xq.hc = client
	xq.token = xq.getToken()

	go func() {
		for {
			time.Sleep(time.Minute * 60)
			xq.token = xq.getToken()
			logrus.Infoln("token updated")
			if xq.token == "" {
				logrus.Warnln("failed to get token")
				return
			}
		}
	}()

	if xq.token == "" {
		logrus.Warnln("failed to get token")
		return nil
	}

	return xq
}

func (x *XueQiu) getToken() string {
	req, err := http.NewRequest("GET", XueQiuHost, nil)
	if err != nil {
		fmt.Println("failed to request xueqiu: ", err)
		return ""
	}

	req.Header.Set("User-Agent", UserAgent)

	resp, err := x.hc.Do(req)
	if err != nil {
		fmt.Println("failed to request xueqiu: ", err)
		return ""
	}
	defer resp.Body.Close()

	for _, cookie := range resp.Cookies() {
		if cookie.Name == TokenName {
			fmt.Printf("Name: %s, Value: %s\n", cookie.Name, cookie.Value)
			return cookie.Value
		}
	}
	return ""
}

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
	req, err := http.NewRequest("GET", qUrl, nil)
	if err != nil {
		return KLine{}, err
	}

	req.Header.Set("User-Agent", UserAgent)
	req.Header.Set("Cookie", fmt.Sprintf("%s=%s;", TokenName, x.token))

	resp, err := x.hc.Do(req)
	if err != nil {
		return KLine{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return KLine{}, fmt.Errorf("failed to get kline: %s", resp.Status)
	}

	var kline KLine
	if err := json.NewDecoder(resp.Body).Decode(&kline); err != nil {
		return KLine{}, err
	}

	return kline, nil
}
