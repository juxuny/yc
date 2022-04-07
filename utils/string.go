package utils

import (
	"github.com/fatih/camelcase"
	"strings"
)

func ToUpperFirst(s string) string {
	ret := []byte(s)
	if len(ret) > 0 && ret[0] >= 'a' && ret[0] <= 'z' {
		ret[0] = ret[0] - ('a' - 'A')
	}
	return string(ret)
}

func ToLowerFirst(s string) string {
	ret := []byte(s)
	if len(ret) > 0 && ret[0] >= 'A' && ret[0] <= 'Z' {
		ret[0] = ret[0] + ('a' - 'A')
	}
	return string(ret)
}

func splitVariable(v string) []string {
	if strings.Contains(v, "_") {
		return strings.Split(v, "_")
	} else if strings.Contains(v, "-") {
		return strings.Split(v, "-")
	} else {
		return camelcase.Split(v)
	}
}

func ToUnderLine(variableName string) string {
	s := splitVariable(variableName)
	for i := range s {
		s[i] = strings.ToLower(s[i])
	}
	return strings.Join(s, "_")
}

func ToHump(variableName string) string {
	s := splitVariable(variableName)
	for i := range s {
		s[i] = strings.ToLower(s[i])
		s[i] = ToUpperFirst(s[i])
	}
	return strings.Join(s, "")
}
