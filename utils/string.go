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
	return strings.Trim(strings.Join(s, "_"), "_")
}

func ToHump(variableName string) string {
	s := splitVariable(variableName)
	for i := range s {
		s[i] = strings.ToLower(s[i])
		s[i] = ToUpperFirst(s[i])
	}
	return strings.Join(s, "")
}

type stringHelper struct{}

var StringHelper = stringHelper{}

func (stringHelper) Filter(data []string, f func(index int, item string) bool) []string {
	ret := make([]string, 0)
	for index, item := range data {
		if f(index, item) {
			ret = append(ret, item)
		}
	}
	return ret
}

func (stringHelper) Transform(data []string, f func(in string) string) []string {
	ret := make([]string, 0)
	for _, item := range data {
		ret = append(ret, f(item))
	}
	return ret
}

func (stringHelper) RandString(n int) string {
	tb := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	ret := make([]byte, 0)
	r := getRandInstance()
	for i := 0; i < n; i++ {
		ret = append(ret, tb[r.Intn(len(tb))])
	}
	return string(ret)
}

func (stringHelper) RandNumString(n int) string {
	tb := "0123456789"
	ret := make([]byte, 0)
	r := getRandInstance()
	for i := 0; i < n; i++ {
		ret = append(ret, tb[r.Intn(len(tb))])
	}
	return string(ret)
}

func (stringHelper) Reverse(input string) string {
	ret := make([]byte, len(input))
	for i, c := range []byte(input) {
		ret[len(input)-1-i] = c
	}
	return string(ret)
}

func (stringHelper) TrimSubSequenceLeft(input, sub string) string {
	for strings.Index(input, sub) == 0 {
		input = strings.Replace(input, sub, "", 1)
	}
	return input
}

func (stringHelper) TrimSubSequenceRight(input, sub string) string {
	reverseInput := StringHelper.Reverse(input)
	reverseSub := StringHelper.Reverse(sub)
	for strings.Index(reverseInput, reverseSub) == 0 {
		reverseInput = strings.Replace(reverseInput, reverseSub, "", 1)
	}
	return StringHelper.Reverse(reverseInput)
}

func (stringHelper) ContainsAllKey(s string, keys []string) bool {
	count := 0
	for _, k := range keys {
		if strings.Contains(s, k) {
			count += 1
		}
	}
	return count == len(keys)
}

func (stringHelper) LowerFirstHump(in string) string {
	return ToLowerFirst(ToHump(in))
}

func (stringHelper) UpperFirstHump(in string) string {
	return ToUpperFirst(ToHump(in))
}
