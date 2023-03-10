package utilsValidator

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"

	"github.com/parinyapt/PT-Friendship_Backend/model/utils"

	"github.com/parinyapt/PT-Friendship_Backend/utils"
	"github.com/parinyapt/PT-Friendship_Backend/utils/response"
)

func ApiInputValidator(c *gin.Context, validateStruct interface{}) (validatePass bool) {
	validate := validator.New()
	if err := validate.Struct(validateStruct); err != nil {
		var listValidateError []modelUtils.ValidatorErrorFieldListStruct
		for _, err := range err.(validator.ValidationErrors) {
			jsonfieldname, errGetStructTagValue := utils.GetStructTagValue(modelUtils.GetStructTagValueParam{
				SelectStruct: validateStruct,
				FieldName:    err.Field(),
				TagName:      "json",
			})
			if errGetStructTagValue != nil {
				utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
					ResponseCode: 500,
				})
				log.Printf("[ Error Report ]->Validate Input Fail | [ Error Detail ]->%s", errors.Wrap(errGetStructTagValue, "[ ApiInputValidator() ]->Get Json Field Name Error"))
				return false
			}
			listValidateError = append(listValidateError, modelUtils.ValidatorErrorFieldListStruct{
				Field:    jsonfieldname,
				ErrorMsg: GetValidatorErrorMessage(err.Tag()),
			})
		}

		utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
			ResponseCode: 400,
			Message:   "Invalid Input Format",
			ErrorCode: "IIF01",
			Data:      listValidateError,
		})
		return false
	}
	return true
}
