package redis

import (
	"github.com/juxuny/yc/utils"
	"reflect"
	"strings"
)

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

func checkIfAllowedData(pkgPath, typeName string) bool {
	return pkgPath == "github.com/juxuny/yc/redis" && typeName == "redis.Key"
}

func InitKeyStruct(v interface{}, upper bool, prefix ...string) {
	vv := reflect.ValueOf(v)
	if vv.Kind() == reflect.Ptr {
		vv = vv.Elem()
	}
	tt := vv.Type()
	for i := 0; i < tt.NumField(); i++ {
		f := tt.Field(i)
		fn := utils.StringHelper.ToUnderLine(f.Name)
		ft := f.Type
		if !checkIfAllowedData(ft.PkgPath(), ft.String()) {
			panic("not allowed data type: " + ft.String() + ", field:" + f.Name)
		}
		if len(prefix) > 0 {
			fn = strings.Join([]string{prefix[0], fn}, "_")
		}
		if upper {
			fn = strings.ToUpper(fn)
		}
		vv.Field(i).SetString(fn)
	}
}

type Config struct {
	UsePass bool
	Host    string
	User    string
	Pass    string
	Port    int
	Index   int
}
