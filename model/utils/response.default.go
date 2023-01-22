package modelUtils

type DefaultResponse struct {
	ResponseCode          int
	DefaultResponseDetail DefaultResponseDetail
}

type DefaultResponseDetail struct {
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
	ErrorCode string      `json:"error_code"`
	Data      interface{} `json:"data"`
}
