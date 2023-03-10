package handler

import (

	"github.com/gin-gonic/gin"

	"github.com/parinyapt/PT-Friendship_Backend/utils/response"
	"github.com/parinyapt/PT-Friendship_Backend/utils/validator"

	"github.com/parinyapt/PT-Friendship_Backend/model/utils"
)

type TestInputStruct struct {
	Email   string `json:"email" validate:"required,email"`
	OTPCode string `json:"otp_code" validate:"required,numeric,len=6"`
	RefID   string `json:"ref_id" validate:"required,uuid"`
}

func DemoHandler(c *gin.Context) {
	var testInput TestInputStruct

	if err := c.ShouldBindJSON(&testInput); err != nil {
		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: 500,
		})
		return
	}
	
	if validatePass := utilsValidator.ApiInputValidator(c, testInput); !validatePass {
		return
	}

	if validatePass := utilsValidator.DataValidator(testInput); !validatePass {
		return
	}
	
}