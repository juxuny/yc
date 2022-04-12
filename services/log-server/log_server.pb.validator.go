package log_server

import "github.com/juxuny/yc/validator"

var templateList = []string{}

func init() {
	if err := validator.RegisterTemplate(templateList...); err != nil {
		panic(err)
	}
}

func (m *PrintRequest) Validate() error {
	return nil
}
