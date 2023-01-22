package utilsResponse

import (
	"github.com/gin-gonic/gin"

	"github.com/parinyapt/PT-Friendship_Backend/model/utils"
)

func ApiDefaultResponse(c *gin.Context, param modelUtils.DefaultResponse) {
	c.JSON(param.ResponseCode, modelUtils.DefaultResponseDetail{
		Success:   param.DefaultResponseDetail.Success,
		Message:   param.DefaultResponseDetail.Message,
		ErrorCode: param.DefaultResponseDetail.ErrorCode,
		Data:      param.DefaultResponseDetail.Data,
	})
}
