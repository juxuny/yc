package validator

import (
	"fmt"
	"github.com/juxuny/yc/errors"
	"github.com/juxuny/yc/utils/template"
)

type IValidator interface {
	Run(value interface{}, refValue string) (bool, error)
}

type Formula string

const (
	FormulaMin       Formula = "min"
	FormulaMax       Formula = "max"
	FormulaIn        Formula = "in"
	FormulaLengthMax Formula = "length.max"
	FormulaLengthMin Formula = "length.min"
)

var validatorSet = map[Formula]IValidator{
	FormulaMin:       &minValidator{},
	FormulaMax:       &maxValidator{},
	FormulaIn:        &inValidator{},
	FormulaLengthMax: &lengthMaxValidator{},
	FormulaLengthMin: &lengthMinValidator{},
}

type Action struct {
	ValidatorFormulas Formula
	RefValue          string
	ErrorTemplate     string
}

func Run(v interface{}, action Action, inputEntity interface{}) error {
	validator, b := validatorSet[action.ValidatorFormulas]
	if !b {
		return errors.SystemError.InvalidValidatorFormula.Wrap(fmt.Errorf("%s", action.ValidatorFormulas))
	}
	if ok, err := validator.Run(v, action.RefValue); err != nil {
		return err
	} else if !ok && action.ErrorTemplate != "" {
		if msg, err := defaultTemplateExecutor.Exec(action.ErrorTemplate, inputEntity); err != nil {
			return err
		} else {
			return errors.SystemError.InvalidParams.SetMsg(msg)
		}
	} else if !ok {
		return errors.SystemError.InvalidParams.WithField("param", v)
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
