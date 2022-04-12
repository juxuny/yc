package validator

import (
	"fmt"
	"github.com/juxuny/yc/errors"
	"github.com/juxuny/yc/utils/template"
)

type IValidator interface {
	Run(value interface{}, refValue string) (bool, error)
}

var validatorSet = map[string]IValidator{
	"min": &minValidator{},
	"max": &maxValidator{},
}

type Action struct {
	ValidatorFormulas []string
	ErrorTemplate     string
}

func Run(v interface{}, refValue string, action Action, inputEntity interface{}) error {
	for _, a := range action.ValidatorFormulas {
		validator, b := validatorSet[a]
		if !b {
			return errors.SystemError.InvalidValidatorFormula.Wrap(fmt.Errorf("%s", a))
		}
		if ok, err := validator.Run(v, refValue); err != nil {
			return err
		} else if !ok && action.ErrorTemplate != "" {
			if msg, err := defaultTemplateExecutor.Exec(action.ErrorTemplate, inputEntity); err != nil {
				return err
			} else {
				return errors.SystemError.InvalidParams.SetMsg(msg)
			}
		} else if !ok {
			return errors.SystemError.InvalidParams
		}
	}
	return nil
}

var defaultTemplateExecutor = template.NewExecutor()

func RegisterTemplate(text ...string) error {
	for _, item := range text {
		if err := defaultTemplateExecutor.Prepare(item); err != nil {
			return err
		}
	}
	return nil
}
