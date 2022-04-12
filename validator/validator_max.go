package validator

import (
	"fmt"
	"github.com/juxuny/yc/errors"
	"strconv"
)

type maxValidator struct {
}

func (t *maxValidator) Run(v interface{}, refValueString string) (bool, error) {
	inputString := fmt.Sprintf("%v", v)
	inputValue, err := strconv.ParseFloat(inputString, 64)
	if err != nil {
		return false, nil
	}
	refValue, err := strconv.ParseFloat(refValueString, 64)
	if err != nil {
		return false, errors.SystemError.InvalidValidatorErrorTemplate.Wrap(err)
	}
	return inputValue <= refValue, nil
}
