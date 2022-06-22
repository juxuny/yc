package redis

import "testing"

func TestInitKeyStruct(t *testing.T) {
	var KeyValue = struct {
		Db Key
	}{}
	InitKeyStruct(&KeyValue, false)
}
