package env

var DefaultEnv = struct {
	Mode string `env:"prod"`
}{}

func init() {
	Init(&DefaultEnv, true)
}
