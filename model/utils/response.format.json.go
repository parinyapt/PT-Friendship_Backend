package modelUtils

type ApiJsonResponseStruct struct {
	ResponseCode int
	Detail       ApiJsonResponseStructDetail
}

type ApiJsonResponseStructDetail struct {
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
	ErrorCode string      `json:"error_code"`
	Data      interface{} `json:"data"`
}
