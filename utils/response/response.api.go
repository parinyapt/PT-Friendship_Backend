package utilsResponse

import (
	"github.com/gin-gonic/gin"

	"github.com/parinyapt/PT-Friendship_Backend/model/utils"
)

var ApiResponseConfigData = map[int]modelUtils.ApiResponseConfigStruct{
	200: {
		Message: "Success",
		ErrorCode: "0",
	},
	400: {
		Message: "Bad Request",
		ErrorCode: "DF400",
	},
	500: {
		Message: "Internal Server Error",
		ErrorCode: "DF500",
	},
}

func ApiResponse(c *gin.Context, param modelUtils.ApiResponseStruct) {
	if param.ResponseCode == 0 {
		param.ResponseCode = 500
	}

	var success_status bool
	if param.ResponseCode >= 200 && param.ResponseCode <= 299 {
		success_status = true
	} else {
		success_status = false
	}

	if param.Message == "" {
		param.Message = ApiResponseConfigData[param.ResponseCode].Message
	}

	if param.ErrorCode == "" {
		param.ErrorCode = ApiResponseConfigData[param.ResponseCode].ErrorCode
	}

	ApiJsonResponse(c, modelUtils.ApiJsonResponseStruct{
		ResponseCode: param.ResponseCode,
		Detail: modelUtils.ApiJsonResponseStructDetail{
			Success:   success_status,
			Message:   param.Message,
			ErrorCode: param.ErrorCode,
			Data:      param.Data,
		},
	})
}
