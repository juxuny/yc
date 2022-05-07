package validator

import (
	"fmt"
	"strings"
)

type passwordValidator struct {
}

func (t *passwordValidator) Run(v interface{}, refValueString string) (bool, error) {
	const upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const lower = "abcdefghijklmnopqrstuvwxyz"
	const number = "0123456789"
	const special = "$%^&*()!@#,./?<>;':[]{}"
	inputString := fmt.Sprintf("%v", v)
	if strings.Contains(refValueString, "letter") && !strings.ContainsAny(inputString, upper) && !strings.ContainsAny(inputString, lower) {
		return false, nil
	}
	if strings.Contains(refValueString, "up") && !strings.ContainsAny(inputString, upper) {
		return false, nil
	}
	if strings.Contains(refValueString, "low") && !strings.ContainsAny(inputString, lower) {
		return false, nil
	}
	if strings.Contains(refValueString, "num") && !strings.ContainsAny(inputString, number) {
		return false, nil
	}
	if strings.Contains(refValueString, "special") && !strings.ContainsAny(inputString, special) {
		return false, nil
	}
	return true, nil
}
