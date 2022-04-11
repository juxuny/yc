package config

import "github.com/juxuny/yc/env"

var Env = struct {
	LogDir       string
	FlushSeconds int    `env:"30"`
	CacheSize    int    `env:"10000"`
	FilePrefix   string `env:"yc"`
	RpcPort      int    `env:"20443"`
	HttpPort     int    `env:"20080"`
}{}

func init() {
	env.Init(&Env, true)
}
