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
	xq.cookie = `cookiesu=601734525893923; HMACCOUNT=49EE4E54B0ED24FF; device_id=a49d66930766010de858ae6ae4ee166f; s=ce17jnz3jb; snbim_minify=true; bid=3474ff7aa50f5d656a5bea7b6d2a4a13_m7lug4nb; xq_is_login=1; u=8240766375; Hm_lvt_1db88642e346389874251b5a1eded6e3=1748948732; xq_a_token=c2256614cb013e6d90463ebd7301e0e8d759c704; xqat=c2256614cb013e6d90463ebd7301e0e8d759c704; xq_id_token=eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJ1aWQiOjgyNDA3NjYzNzUsImlzcyI6InVjIiwiZXhwIjoxNzUyMTQ4MTg3LCJjdG0iOjE3NDk1NTYxODczNzksImNpZCI6ImQ5ZDBuNEFadXAifQ.q7vgSCkAZj0GWw_C2HxPK8H8UNlJXOP13S1r6w0Y9uIa41e10s960vVXXTvhWHAoA0DyFYZy9wdMz9zqf9CO_aZ1b327nDrfjUegI6jBE9DUYMGtNM7BVIxK8ktJ0bU2W3l2jNhVYWpqv-8J6xqjrfL1Yt_dn1kWG281YxVo1n5cUu8LiCZfdloulA_A0hQNNY1u6COYwUfdnO2hXsKnqZsbrMnaQ8aUVeoNjWsBna03JOc1RX4Z7nFw7nhX-ViHyFLO9to_XSgsO5K-JrUdszp0dh60mBRvzf0foUnKl3WSwaa6hGEZ2R01KUUuPAhUxb94zWQM5j4gBjq4fodOgA; xq_r_token=9d642c1617ce2cb9b0ec5837a2b79449aa9449dd; is_overseas=1; Hm_lpvt_1db88642e346389874251b5a1eded6e3=1749556234; ssxmod_itna=eqmh7KAKGKDI/DUxe9Dmu340OxKwAx79eGHDyxWKG7DuxiK08D6iDBRr1+dgin+a+eDj2GIKDoolDNDl=YxoDSxD=ODK4GTmm60YNhr8mBgDve+7ODo4qkBChqIKt72QhRjOKhGl9X10qoDnne1DibDBKDnh4Dj2TqmDii7x0rD0eDPxDYDG+mDD5A=+4Djjop7REGAABe=UoDbh4+UG4DRhdDSGfAmUEGdC3DahYdG+b5DmFY3w+dk4GaBR+RDlK+Dm4Ta86cDC9Xo9r=zKSAAmZ8dYy4meoYmmGT5QGK7DhiUX01rQixrismGsk048ixFGDx3bQmeDD3bGIY/UxtDhi+Uxt2hYRr4=jIqnvgaIZ0eAYx4Ye2Ye=Y2pnwsee+0NKDpe7xdAXZmDP4NdYD; ssxmod_itna2=eqmh7KAKGKDI/DUxe9Dmu340OxKwAx79eGHDyxWKG7DuxiK08D6iDBRr1+dgin+a+eDj2GIKDoolDoDipjxXs8Dj4xsTP9DDs=It0LK=8V91xx6eI9Z0xnfvz72BXnGyIiHOojL=/WOqj91rlyDsPvF4Z5hj=Fha6Giw10kbAc4W6ICr45seaPCRon4a9F4E1B1ARHHLpGFDFgO1ZpxNDdOh4Mf=f5AZbO8YFvGBFSUz2H12tQ1ErGiF0aR5QG2Ez=ss1BL64yOD4hHoicxO6pQrMCHF7rkScp+xZ3aca1sMgMxvoW/3UXFQRs3yO/2+/97WmZDHUDAQAn87YctqhtnOQmjSTKLYsYPujIyD7UjieI951nowajjboxepzYCmmzCK4phPzF=qnH7mcjEPlRVzvifpP2cH9uZAWraW2bu8DgkRRnSRKnRzzSpf5bX74RwOi7c/EQBD4=qhDmxGwcWD1QWi=fBjRub0xeq=+iPkDhBidk7NupqYr4+Gzgx64qaYiiIRGcMCDT7fR+2istCmalzKGgKq69yVe4wj0DzPKcQn+gzo1HnG+v+OYc4QqEV4LyuLv+K+DoYxWS=oFbN/riY+jHi4B=RHDoYDXPD`

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
