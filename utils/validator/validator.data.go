package utilsValidator

import (
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

func DataValidator(validateStruct interface{}) (validatePass bool) {
	validate := validator.New()
	if err := validate.Struct(validateStruct); err != nil {
		validateError := errors.New("[ DataValidator() ]->Data Invalid Format")
		for _, err := range err.(validator.ValidationErrors) {
			validateError = errors.Wrap(validateError, fmt.Sprintf("[ Error Field ]->%s->%s", err.Field(), GetValidatorErrorMessage(err.Tag())))
		}
		log.Printf("[ Error Report ]->Validate Data Invalid Format | [ Error Detail ]->%s", validateError)
		return false
	}
	return true
}