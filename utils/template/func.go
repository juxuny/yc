package template

func toUpperFirst(s string) string {
	ret := []byte(s)
	if len(ret) > 0 && ret[0] >= 'a' && ret[0] <= 'z' {
		ret[0] = ret[0] - ('a' - 'A')
	}
	return string(ret)
}

func toLowerFirst(s string) string {
	ret := []byte(s)
	if len(ret) > 0 && ret[0] >= 'A' && ret[0] <= 'Z' {
		ret[0] = ret[0] + ('a' - 'A')
	}
	return string(ret)
}
