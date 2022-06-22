package config

import (
	"github.com/juxuny/yc/redis"
	"github.com/juxuny/yc/services/cos"
)

type keyValue struct {
	Db            redis.Key
	NotifyChannel redis.Key
}

var KeyValue keyValue

func init() {
	redis.InitKeyStruct(&KeyValue, false, cos.Name)
}
