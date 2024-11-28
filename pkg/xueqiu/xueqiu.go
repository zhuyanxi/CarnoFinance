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
	xq.cookie = `cookiesu=901709550598480; device_id=f19d6e2976d8f0a2558fb5aa5bd6ebdd; snbim_minify=true; s=c911rc7q1i; bid=3474ff7aa50f5d656a5bea7b6d2a4a13_lvu851fh; HMACCOUNT=00A27D732039D4B8; xq_is_login=1; u=8240766375; Hm_lvt_1db88642e346389874251b5a1eded6e3=1730514979; xq_a_token=38bc741aaaf73162d37ea4833ebaafc1dea559e4; xqat=38bc741aaaf73162d37ea4833ebaafc1dea559e4; xq_id_token=eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJ1aWQiOjgyNDA3NjYzNzUsImlzcyI6InVjIiwiZXhwIjoxNzM0NDE0ODEzLCJjdG0iOjE3MzE4MjI4MTMxODUsImNpZCI6ImQ5ZDBuNEFadXAifQ.pHov6RwzrKolNfgV4BsS8x3f5dgiPeEX5W4Gdm5Qn_EL3Umui5Kd2vKKTgTKKTodztBXbPzjzHjgK-FO7VgqvsrA3iv2-03s_3yZ_Y81wucrLbOJ3kU2dBbTdXOfXq7LxYaGhEgEDpIEwkSFkvdJqfG3-F2n44qZcZiVrd3n6UuzvT7UWzxUGV_pXLslOi9kQAbm8amF5YSl8a3p80plt_ohBe40_44rWKlDOHXcXzavs_d3TSnFBSzkVTsLPix_o-KUk_kD459qMW3YQBHOZz5v52cHIM9lTPOFPRRK_IthqYbGP0wbbwThnSC6-PcA5Ofc553_dQ7gqocLpWA8gw; xq_r_token=3cbcdc0543691dfe056cef6dbf9bc40e637a502e; is_overseas=0; Hm_lpvt_1db88642e346389874251b5a1eded6e3=1731936191; ssxmod_itna=eqIO0KAKDKGKeGq0dGQDHFFfCDU2Dmq22DmOH+eq5DsqEDSxGKidDqxBme1qnbc43QbmGPq4CDrOhLP/Cxo2DFitD+PP4j/=i4GIDeKG2DmeDyDi5GRD0ImKbDeetD5xGoDPxDeDAGqGamxDrmPvz9D0kDY5D06Y7Axi8Y8KDjA5GqFAznEKDK6A7f2oDWE53dBpzDlewfO91ShvODCFmIE9p6/8j=72DDm4dnhY+nR9AvoIPZpohKI7wN3opDHrq7xyNWQrrPW4Y=j0Y0rhbqixPqj0xpvPx=maqGQrHVCNDic0xbGnD4D=; ssxmod_itna2=eqIO0KAKDKGKeGq0dGQDHFFfCDU2Dmq22DmOH+eqikL=W5DlabjDj4a5mZLQidqjgLw=Ye+gwDPhqFjeePm=YYm0UerxIBDXK9SQYrj1/Ft7DYEQ6=Gb9ejOV3KFzWqvD4bu9QQYKaDZp7yY7=2jWOZciQPbpbYPeidrDw/DHO0r0rDoOtd+eHT9HFz9v7ohqkGeFm=mGH=x5q+ZffeY+WerxaGjPGy7k5djDbexnWj60xVBxCAx4tiY9D2bgskcyti9WnQiSQ5iPQZ1W7eP6x2AqNGFisZaQsktKTyDoIs4ENmOG9rFPUiNmoek6XDcmRcYkCweiAkPpgPpMDU9DNUP=XAmaDG23xPeT=7T4Dxrl01GDYv2=evKN0=Gi5n2Q4bdMuddR4D08DijPYD=`

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
