package xueqiu

import (
	"fmt"
	"io"
	"net/http"
)

const (
	UserId     = "8240766375"
	XueQiuHost = "https://xueqiu.com/today"
	UserAgent  = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36"
	TokenName  = "xq_a_token"
)

type XueQiu struct {
	hc *http.Client

	// token string

	cookie string
}

func New() *XueQiu {
	xq := new(XueQiu)
	client := &http.Client{}

	xq.hc = client
	// xq.token = xq.getToken()
	xq.cookie = `cookiesu=901709550598480; device_id=f19d6e2976d8f0a2558fb5aa5bd6ebdd; snbim_minify=true; s=c911rc7q1i; bid=3474ff7aa50f5d656a5bea7b6d2a4a13_lvu851fh; HMACCOUNT=00A27D732039D4B8; xq_is_login=1; u=8240766375; Hm_lvt_1db88642e346389874251b5a1eded6e3=1726658449; xq_a_token=337abe5b962be3520c28543c6fb1c77c98c5e4b8; xqat=337abe5b962be3520c28543c6fb1c77c98c5e4b8; xq_id_token=eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJ1aWQiOjgyNDA3NjYzNzUsImlzcyI6InVjIiwiZXhwIjoxNzMwODYxNjUxLCJjdG0iOjE3MjgyNjk2NTE0MjYsImNpZCI6ImQ5ZDBuNEFadXAifQ.GvQj1uml55QLvvag8X9uSBe-AzAFZL9o4Q1Ajfza7NyWaOWTOP9F7OXajPpMHRdYwm_QnwdsG87-mDtfWJTxelsIf6M-0sxd__ihh1ak1Honw6sy8IJc59f0w8NYeEvX3hKJQ5Pu-fyj_lHo5lg2sfxB0nHGFKThiYrHEakfYtfUSlZtYqH2Fm5fFc5LfeY8ATT0a4WBKGs9nETzNG0vdeRE7LpU-HnRHUo6du0d19pFr57GO957AQFwtInFuJJMi39NZ30-tJiJvK3rv3Ytt8qf2TpVgCJIDuoPyDaH2SsicGg7MraYDphCtdh4M8fmssKfhzGMqSqP04RaN0fWbg; xq_r_token=f05a34d942aa8fcee548fe4b3c78258575e8662d; is_overseas=0; Hm_lpvt_1db88642e346389874251b5a1eded6e3=1728389340; ssxmod_itna=YqUxgD9D0Du0KAKP0Lx0Pp1Qx2iKDB0D8BDzDFxRcx055neGzDAxn40iDt=OqiqY3qWQi4eY6m0AwhhXWelcfx6Dh0iLxQzcO4GIDeKG2DmeDyDi5GRD09esTDemtD5xGoDPxDeDAAqGamxDrmPvzBD0kDY5D06Y7Axi8Y8KDjA5GqFAznEKDK6A7f2oDWE5xbBpzDlewfO91ShvODCFmIE9p6/8j/7x5DmbdnhY+nR9AvoIPZpohKI7wN3opDHhNOxviEGe5Q0xo/nD40=iTCQ0DCADgaAPq/ApNKm4djih4DDcS9kXDD==; ssxmod_itna2=YqUxgD9D0Du0KAKP0Lx0Pp1Qx2iKDB0D8BDzDFxRxnKMDmxDseC3DLBxpLi4oqAp7yDbBKwxR0hqxxpgtRBovwDArKaft0zq5bU8VOMFcAR3CqhaOaEFzkuCKyFS9DqL=GX3qu9SNeqH2x4qmI8sgrszjK5in5iwCWbe7eiYAcaDOYiQ37itbnk8Rxz1toAMCSvk9p13QqmwcFiiRiIcbWYLc8vupnEr9Y5CtiFKVq4pcoYM0D4bcOQK/l+a=iNxrisWLibKSabOh8spV=aBZPeKbYfK7nDWdtetq4fz7N8krQy52hQGD5k+8tET4ek5rxcRbhSFxE39LTXx7iGhEq7YiaC0zQihDhNeh7A4YIROorG8gmUf33eCARpsxIsuatjQWT3poCvx3sOoyD4YBoQxAGRakWTZrv4hIDdW3cY3cpeRwR8moPa/mNu0vkgfo4bIWK2OGyqA+3jy+TShCOQmklYGDW1eNqA3KhqYhWjQ=BBD7xF++QqDvQZQB2dVTenW3CA3Ib=bBQarwTlXFS=2c=wjWyB+xgtgYpszN/qWycQPE0xuCPZwpt7yjcRdHL47CbnxTdVdyktmrhPD7QNA5On2CBBlhDM=OiCPwMAP7v50HP/CS6wBh5oWZjiNkUL7hq5xxfeXRg0kpqYDrM0gVOZc+61iAYlys+mwQm=3YKTp5r6hGAKqQEC0me3rA44lYQipxDFqDeTxeD==`

	// go func() {
	// 	for {
	// 		time.Sleep(time.Minute * 60)
	// 		xq.token = xq.getToken()
	// 		logrus.Infoln("token updated")
	// 		if xq.token == "" {
	// 			logrus.Warnln("failed to get token")
	// 			return
	// 		}
	// 	}
	// }()

	// if xq.token == "" {
	// 	logrus.Warnln("failed to get token")
	// 	return nil
	// }

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

	cookies := resp.Cookies()

	for _, cookie := range cookies {
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
	// req.Header.Set("Cookie", fmt.Sprintf("u=%s;%s=%s;", UserId, TokenName, x.token))
	req.Header.Set("Cookie", x.cookie)

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
