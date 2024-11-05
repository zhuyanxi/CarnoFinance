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
	xq.cookie = `cookiesu=901709550598480; device_id=f19d6e2976d8f0a2558fb5aa5bd6ebdd; snbim_minify=true; s=c911rc7q1i; bid=3474ff7aa50f5d656a5bea7b6d2a4a13_lvu851fh; HMACCOUNT=00A27D732039D4B8; xq_is_login=1; u=8240766375; xq_a_token=a6fad56af39a9b23ff9c84ea871e91d7b582c1f6; xqat=a6fad56af39a9b23ff9c84ea871e91d7b582c1f6; xq_id_token=eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJ1aWQiOjgyNDA3NjYzNzUsImlzcyI6InVjIiwiZXhwIjoxNzMzMDY2NTE3LCJjdG0iOjE3MzA0NzQ1MTc4NjMsImNpZCI6ImQ5ZDBuNEFadXAifQ.me6H2qy2Ts8jF_AZtMk971LveZjNLDDgwb6El1ZmoioBD5gmYlU8JC4Ka6xXTMxLwJ8GZLzyJ8aZKMK906l16tQJLmSRfuhf8J-JCe7OZFF2JrDOgdRGpBC60LYDzGZQk_vp1EXgHwynpf-jk1X3w5-JA5zV76CZzB1EomCD4AS4yHG1kIh0SQfiHfXTdvxTaibWLrODcdkka4kWdWuq-bxMAuSDJHIqFRCa9ZiXoTGAh8vxxlE-jOI0IsKUHd9CQpeOc8Rn3kzfJvBRd66riuXO8Yc9rJwQHXM5kqxOx-RZnUSYqW7IoQOpr5XAOz6qWAbJDhghhZXKS8D5qjcs5w; xq_r_token=6f8d2ae583d04f9323662332a7711136299d0849; Hm_lvt_1db88642e346389874251b5a1eded6e3=1730514979; Hm_lpvt_1db88642e346389874251b5a1eded6e3=1730625886; is_overseas=0; ssxmod_itna=YuG=DvqUxAhhO7DXKDHD0W5biQfypRDitqCCQ3wb5Dcnx05ydGzDAxn40iDtohuB0P3HKnSbLF0t/ATEdRUYWIdt0tW7ejxNG=mweD=xYQDwxYoDUxGtDpxG=oTDeeND5xGoDPxDeDA0qDF1KDe19ClEqDXxGCDDHww6KDEE78D04s4FBhlOHkD7WwpcIvDiPs3CW7lx0Tt9F8g2vzEx0OTIFkO/8juRiPCDi307v+38MkpfFOalyIK+B+NKopF3QPQjceq=0xbAe5N02xw00PNW0em74xCChSW0aeC20Ysq9iSqDW5DFTCiDD==; ssxmod_itna2=YuG=DvqUxAhhO7DXKDHD0W5biQfypRDitqCCQ3wb5DcDn9M0q5DsCrwDLQCa0TL9qPSlAM2Yhe21mr=aPyZBr6qV7tZgUF=Lc+MpwqfcT5GFBn10gnOT5UG+8hrycl8uNtb=NMUhqLh8SMt1G70c4dL4dQdSbM05N3ES4whdiW5dkO0nfMh23WrtPajxOiuwYaX9BqaPgwLF5WN7vdcPsacxV29MUBG+8g4Wg7mdb8eU/=QtKiUc8kjy9cuOcL043BIu9hw4BqP87qWG0Euh3+e13Gh6VO=ifhhniEL4lF3vbBAw0le+AAE=wu0Pi1wd7=WZG6OpuD+wj65jQ4j5RG5w1ekBQ9Y6vCK+mPjPNfGvBhoT83f1+E2qdUIBOq7zegmIB2E5Z=UGDQ1vpOmq8IFGrN3W07fMi6GR3mUIiK3FOifeIvOfAOik0pB=Iwz0UtIWzio1ip4FC=ID8c+cP1x+HDG2B0HA4eV9YipxfyCrYMUmfB+Yt0VbTMKpjCKDFqD+1xxD`

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
