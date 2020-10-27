package vo

const (
	ErrorCodeSuccess = 0
	SuccessStr       = "Success"
)

type ResultBase struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	UserMsg string `json:"userMsg"`
}

type Result struct {
	ResultBase
	Data interface{} `json:"data"`
}

type StringResult struct {
	ResultBase
	Data string `json:"data"`
}

func NewSuccessResult(data interface{}) *Result {
	r := new(Result)
	r.Code = ErrorCodeSuccess
	r.Msg = SuccessStr
	r.UserMsg = SuccessStr
	r.Data = data
	return r
}

func NewSuccessStringResult(data string) *Result {
	r := new(Result)
	r.Code = ErrorCodeSuccess
	r.Msg = SuccessStr
	r.UserMsg = SuccessStr
	r.Data = data
	return r
}

func NewUserResult(code int, msg, userMsg string, data interface{}) *Result {
	r := new(Result)
	r.Code = code
	r.Msg = msg
	r.UserMsg = userMsg
	r.Data = data
	return r
}

func NewStringResult(code int, msg, userMsg, data string) *StringResult {
	r := new(StringResult)
	r.Code = code
	r.Msg = msg
	r.UserMsg = userMsg
	r.Data = data
	return r
}
