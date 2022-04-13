package validator

import (
	"fmt"
	"time"
)

type timeValidator struct {
	layout string
}

func (t *timeValidator) Run(v interface{}, refValueString string) (bool, error) {
	inputString := fmt.Sprintf("%v", v)
	layout := t.layout
	if refValueString != "" {
		layout = refValueString
	}
	_, err := time.ParseInLocation(layout, inputString, time.Local)
	if err != nil {
		return false, nil
	}
	return true, nil
}
