package utils

import (
	"reflect"

	"github.com/pkg/errors"
	
	"github.com/parinyapt/PT-Friendship_Backend/model/utils"
)

func GetStructTag(config modelUtils.GetStructTagParam) (string, error) {
	field, ok := reflect.TypeOf(config.SelectStruct).FieldByName(config.FieldName)
	if !ok {
		return "", errors.Wrap(errors.New("[Custom]->Error=False"), "[Error]->Reflect Field Error")
	}

	return string(field.Tag.Get(config.TagName)), nil
}
