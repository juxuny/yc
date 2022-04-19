package orm

const DefaultName = "default"

type Config struct {
	Name   string `json:"name"`
	Host   string `json:"host"`
	Port   int    `json:"port"`
	User   string `json:"user"`
	Pass   string `json:"pass"`
	Schema string `json:"schema"`
}

func InitConfig(config ...Config) error {
	for _, c := range config {
		if err := connectManagerInstance.Add(c); err != nil {
			return err
		}
	}
	return nil
}
