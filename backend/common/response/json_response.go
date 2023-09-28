package response

type Code string

const (
	SuccessCode Code = "1"
	UnknownCode Code = "-1"
)

type JsonReponse struct {
	Code    Code        `json:"code"`
	Msg     string      `json:"msg"`
	Content interface{} `json:"content, omitempty"`
}
