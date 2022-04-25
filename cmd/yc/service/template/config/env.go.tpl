package config

import "github.com/juxuny/yc/env"

var Env = struct {
	RpcPort      int    `env:"20443"`
	HttpPort     int    `env:"20080"`
}{}

func init() {
	env.Init(&Env, true)
}
