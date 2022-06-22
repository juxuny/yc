package config

import (
	"github.com/juxuny/yc/env"
	"github.com/juxuny/yc/jwt"
	"github.com/juxuny/yc/log"
	"github.com/juxuny/yc/orm"
	"github.com/juxuny/yc/redis"
	"github.com/juxuny/yc/services/cos"
	"os"
)

var Env = struct {
	RpcPort  int `env:"20443"`
	HttpPort int `env:"20080"`

	DbUser          string
	DbPass          string
	DbHost          string
	DbPort          int
	DbName          string
	JwtSecret       string
	IgnoreCallLevel bool `env:"true"`
	RedisHost       string
	RedisPort       int
	RedisUser       string
	RedisPass       string
	RedisIndex      int
}{}

func init() {
	env.Init(&Env, true)

	// config database
	c := orm.Config{
		Name:   cos.Name,
		Host:   Env.DbHost,
		Port:   Env.DbPort,
		User:   Env.DbUser,
		Pass:   Env.DbPass,
		Schema: Env.DbName,
	}
	if err := orm.InitConfig(c); err != nil {
		log.Error(err)
		os.Exit(-1)
	}

	// config redis
	redis.InitConfig(redis.Config{
		UsePass: true,
		Host:    Env.RedisHost,
		User:    Env.RedisUser,
		Pass:    Env.RedisPass,
		Port:    Env.RedisPort,
		Index:   Env.RedisIndex,
	})

	// config jwt secret
	jwt.SetSecret(Env.JwtSecret)

}
