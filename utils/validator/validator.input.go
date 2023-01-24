package utilsValidator

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"

	"github.com/parinyapt/PT-Friendship_Backend/model/utils"

	"github.com/parinyapt/PT-Friendship_Backend/utils"
	"github.com/parinyapt/PT-Friendship_Backend/utils/response"
)

func ApiInputValidator(c *gin.Context, validateStruct interface{}) (bool, error) {
	validate := validator.New()
	if err := validate.Struct(validateStruct); err != nil {
		var listValidateError []modelUtils.ValidatorErrorFieldListStruct
		for _, err := range err.(validator.ValidationErrors) {
			jsonfieldname, errGetStructTag := utils.GetStructTagValue(modelUtils.GetStructTagValueParam{
				SelectStruct: validateStruct,
				FieldName:    err.Field(),
				TagName:      "json",
			})
			if errGetStructTag != nil {
				utilsResponse.ApiResponse(c, modelUtils.ApiResponseStruct{
					ResponseCode: 500,
				})
				return false, errors.Wrap(errGetStructTag, "[ ApiInputValidator() ]->Get Json Field Name Error")
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
		return false, nil
	}
	return true, nil
}
