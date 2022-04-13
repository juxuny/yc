package log_server

import (
	"github.com/juxuny/yc/validator"
)

const (
	ValidatorTemplatePrintRequestLevel     = "invalid level"
	ValidatorTemplatePrintRequestContent   = "content cannot be empty"
	ValidatorTimestampPrintRequestFileLine = "fileLine cannot be empty"
)

var templateList = []string{
	ValidatorTemplatePrintRequestLevel,
	ValidatorTemplatePrintRequestContent,
	ValidatorTimestampPrintRequestFileLine,
}

func init() {
	if err := validator.RegisterTemplate(templateList...); err != nil {
		panic(err)
	}
}

func (m *PrintRequest) Validate() error {
	if err := validator.Run(m.Level, validator.CreateAction("in", "1,2,3,4", ValidatorTemplatePrintRequestLevel), m, "level"); err != nil {
		return err
	}
	if err := validator.Run(m.Content, validator.CreateAction("length.min", "1", ValidatorTemplatePrintRequestContent), m, "content"); err != nil {
		return err
	}
	if err := validator.Run(m.DateTime, validator.CreateAction("timestamp.log", "", ""), m, "dateTime"); err != nil {
		return err
	}
	if err := validator.Run(m.FileLine, validator.CreateAction("length.min", "1", ValidatorTimestampPrintRequestFileLine), m, "fileLine"); err != nil {
		return err
	}
	return nil
}
