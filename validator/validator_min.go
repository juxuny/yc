package validator

import (
	"github.com/juxuny/yc/errors"
	"github.com/juxuny/yc/utils"
	"strconv"
)

type minValidator struct {
}

func (t *minValidator) Run(v interface{}, refValueString string) (bool, error) {
	inputValue, err := utils.Converter.ToFloat64(v)
	if err != nil {
		return false, nil
	}
	refValue, err := strconv.ParseFloat(refValueString, 64)
	if err != nil {
		return false, errors.SystemError.InvalidRefValueDefinition.Wrap(err)
	}
	return inputValue >= refValue, nil
}
