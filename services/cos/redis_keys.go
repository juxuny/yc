package cos

import (
	"github.com/juxuny/yc/redis"
)

type keyValue struct {
	NotifyChannel redis.Key
}

var RedisKey keyValue

func init() {
	redis.InitKeyStruct(&RedisKey, false, Name)
}
