package redis

import "strings"

type Key string

func (t Key) String() string {
	return string(t)
}

func (t Key) Suffix(v ...string) Key {
	if len(v) == 0 {
		return t
	}
	return Key(strings.Join(v, "_") + "_" + string(t))
}

func (t Key) Prefix(v ...string) Key {
	if len(v) == 0 {
		return t
	}
	return Key(string(t) + "_" + strings.Join(v, "_"))
}

func convertKeysToStrings(keys ...Key) []string {
	ret := make([]string, len(keys))
	for i, item := range keys {
		ret[i] = item.String()
	}
	return ret
}

type Config struct {
	UsePass bool
	Host    string
	User    string
	Pass    string
	Port    int
	Index   int
}
