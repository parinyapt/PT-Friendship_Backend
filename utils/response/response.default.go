package utilsResponse

import (
	"github.com/gin-gonic/gin"

	"github.com/parinyapt/PT-Friendship_Backend/model/utils"
)

func ApiDefaultResponse(c *gin.Context, param modelUtils.DefaultResponse) {
	c.JSON(param.ResponseCode, modelUtils.DefaultResponseDetail{
		Success:   param.Detail.Success,
		Message:   param.Detail.Message,
		ErrorCode: param.Detail.ErrorCode,
		Data:      param.Detail.Data,
	})
}
