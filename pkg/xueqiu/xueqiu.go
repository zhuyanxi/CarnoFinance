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
	xq.cookie = `cookiesu=901709550598480; device_id=f19d6e2976d8f0a2558fb5aa5bd6ebdd; snbim_minify=true; s=c911rc7q1i; bid=3474ff7aa50f5d656a5bea7b6d2a4a13_lvu851fh; HMACCOUNT=00A27D732039D4B8; remember=1; xq_a_token=597ba7f95349b7efb913fc3a5c63376a63911024; xqat=597ba7f95349b7efb913fc3a5c63376a63911024; xq_id_token=eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJ1aWQiOjgyNDA3NjYzNzUsImlzcyI6InVjIiwiZXhwIjoxNzI3OTU4OTM5LCJjdG0iOjE3MjU2MzEzMDk5NjYsImNpZCI6ImQ5ZDBuNEFadXAifQ.CMj6weIcEdY4F0_jA208dgMYrlVIuLiaxjNxVj6v_Uq4ne_q-kB3pUS5HkPCQRPL2Vun86P4pP-vDFpp80S_lMOKvv8KjdrlT5kScrz3JDxJZPKipgN_E8TH6RE3tTxpr9XB6eUzj0sERunm9dhAV6E7CaOnRFmK3zY9_smA8mnhSFfsptocZ9skiBXwA5znf5F1_Rb7BXFEaMVBh--qeBbQyQqtV9fETcOeKEPWlSelmGUFwVfb38ooJWPpHmPlKwjJnOtv4Q3ZobRQnlY8XM8-VshnZzWNJkBndYej0xAMZ9Plentql04lkhB-0THK46gcI6CaiYBOqKdZFfZp5g; xq_r_token=1e94361a027c9095f875d4a3ecd561e02b399705; xq_is_login=1; u=8240766375; is_overseas=1; Hm_lvt_1db88642e346389874251b5a1eded6e3=1726658449; Hm_lpvt_1db88642e346389874251b5a1eded6e3=1726658492; ssxmod_itna=WqRxuiD=GQw=DXFDHD0QHiiQGkFDk/KD0oOmuMkTdOqGN6YDZDiqAPGhDC4R9jbbpKaKQ0cG3K9i7rP7aKgjqhNUKZhGLPKCxWqeD=xYQDwxYoDUxGtDpxG=Y6DemtD5xGoDPxDeDAAqDLQqD+QTUZ0xGyDitDGwr2WqGnbe=D7OXDon0Zfw=DQ+rR7LTDA8Nti4jZDBEKOq=l9dNnDBQorw=Q1PFOT74=DYPHfd2eLc=bTwjolTu4NWr5hCR3qF+enfPGF7eYQbqqn7ePbSwFFGo=F7x=fBCo9RDKIvDXV4AlxD3yV34=eD; ssxmod_itna2=WqRxuiD=GQw=DXFDHD0QHiiQGkFDk/KD0oOmuMkTdxA69IND/nY/DF27lFqOiKid8FCkxhCK49B9GeYaUwqtqo4KZAhm1UggO9+8Raay63NcjTBgkDobaIg9v50FTBjRYdm/MxC7wabC7DkUaPxhEKzejwOKjD=aCmwVWDNN6DrMlY=PpuN4O9WoxnwlQ8ioe94tFaEe=nT61mhxZiORfhOyg87t4E8RhHEjrBkxuHa2gIFjpQCjgevdVon4vWCy2hcycj+a0ZTI6kogZ5iTxEKqGTkKVwrEOozK1/DtfaTPuLQsCtdgi1trr=wb9ft7gqGdTxlYqra3OxEOxDl7e3Q+YYnr=qi1BbqhiKehO+eHXhYBLTEWYdYxIOiG2Kxr3gD5SRD8iKoBC7ov9gY4xdi3TexYoGNqcerkTfMqP=FCSDIW1roGpW8DuHyD+8ZeR7wVhTKA1bnfoA1KBH+CGaebwLYa7e8lbCRATaon1Y3g+Kq7+YGFPYCBtk+WCi8k3vAk5fbF=t6x9YlBt3KxgY0wZU6gNfPYEgtDLAbwWkW83UAZ=arTF+t0o4LxgcgBncOLGd7UonP3G6DDwZGxlqYqb3z3yelsALoLCnjqbjt8eKM0lrjq1aPpL0EA0kQ4VeZkAZ8UdMDGcDG7aiDD`

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
