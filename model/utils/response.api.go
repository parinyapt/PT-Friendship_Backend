package modelUtils

type ApiResponseStruct struct {
	ResponseCode int
	Message      string
	ErrorCode    string
	Data         interface{}
}

type ApiResponseConfigStruct struct {
	Message   string
	ErrorCode string
}
