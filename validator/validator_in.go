package validator

import (
	"fmt"
	"github.com/juxuny/yc/utils"
	"strings"
)

type inValidator struct {
}

func (t *inValidator) Run(v interface{}, refValueString string) (bool, error) {
	inputString := fmt.Sprintf("%v", v)
	inputString = strings.TrimSpace(inputString)
	items := strings.Split(refValueString, ",")
	items = utils.StringHelper.Transform(items, strings.TrimSpace)
	refSet := utils.NewStringSet(items...)
	return refSet.Exists(inputString), nil
}
