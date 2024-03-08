package xueqiu

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	XueQiuHost = "https://xueqiu.com/"
	UserAgent  = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/000000000 Safari/537.36"
	TokenName  = "xq_a_token"
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

func (x *XueQiu) requestXueqiu(qUrl string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, qUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", UserAgent)
	req.Header.Set("Cookie", fmt.Sprintf("%s=%s;", TokenName, x.token))

	resp, err := x.hc.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed to request xueqiu: %s", resp.Status)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bodyBytes, nil
}
