package tushare

type Tushare struct {
	Token string
}

func NewTushare(token string) *Tushare {
	return &Tushare{
		Token: token,
	}
}

type TushareRequest struct {
	APIName string `json:"api_name,omitempty"`
	Token   string `json:"token,omitempty"`
	Params  any    `json:"params,omitempty"`
	Fields  string `json:"fields,omitempty"`
}

type TushareResponse struct {
	RequestId string              `json:"request_id,omitempty"`
	Code      int                 `json:"code,omitempty"`
	Msg       string              `json:"msg,omitempty"`
	Data      TushareResponseData `json:"data,omitempty"`
}

type TushareResponseData struct {
	Fields []string `json:"fields,omitempty"`
	Items  [][]any  `json:"items,omitempty"`
}
