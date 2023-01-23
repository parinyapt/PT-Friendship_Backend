package utilsResponse

import (
	"github.com/gin-gonic/gin"

	"github.com/parinyapt/PT-Friendship_Backend/model/utils"
)

func ApiJsonResponse(c *gin.Context, param modelUtils.ApiJsonResponseStruct) {
	c.JSON(param.ResponseCode, modelUtils.ApiJsonResponseStructDetail{
		Success:   param.Detail.Success,
		Message:   param.Detail.Message,
		ErrorCode: param.Detail.ErrorCode,
		Data:      param.Detail.Data,
	})
}
