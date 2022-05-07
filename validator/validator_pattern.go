package validator

import (
	"fmt"
	"github.com/juxuny/yc/errors"
	"regexp"
)

type patternValidator struct {
}

func (t *patternValidator) Run(v interface{}, refValueString string) (bool, error) {
	inputString := fmt.Sprintf("%v", v)
	matched, err := regexp.MatchString(refValueString, inputString)
	if err != nil {
		return false, errors.SystemError.InvalidRefValueDefinition.Wrap(err).WithField("pattern", refValueString)
	}
	return matched, nil
}
