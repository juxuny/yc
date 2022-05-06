package config

import (
	"github.com/juxuny/yc/env"
	"github.com/juxuny/yc/jwt"
	"github.com/juxuny/yc/log"
	"github.com/juxuny/yc/orm"
	"github.com/juxuny/yc/services/cos"
	"os"
)

var Env = struct {
	RpcPort  int `env:"20443"`
	HttpPort int `env:"20080"`

	DbUser    string
	DbPass    string
	DbHost    string
	DbPort    int
	DbName    string
	JwtSecret string
}{}

func init() {
	env.Init(&Env, true)

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

	jwt.SetSecret(Env.JwtSecret)
}
