package utils

import (
	"reflect"

	"github.com/pkg/errors"
	
	"github.com/parinyapt/PT-Friendship_Backend/model/utils"
)

func GetStructTagValue(config modelUtils.GetStructTagValueParam) (string, error) {
	field, ok := reflect.TypeOf(config.SelectStruct).FieldByName(config.FieldName)
	if !ok {
		return "", errors.New("[ GetStructTagValue() ]->Reflect Field Error")
	}

	return string(field.Tag.Get(config.TagName)), nil
}
