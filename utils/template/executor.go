package template

import (
	"bytes"
	"fmt"
	"github.com/juxuny/yc/errors"
	"html/template"
)

type Executor interface {
	Prepare(templateData string) error
	Exec(templateData string, data interface{}) (string, error)
	Exists(templateData string) bool
}

type executor struct {
	m map[string]*template.Template
}

func (t *executor) Exists(templateData string) bool {
	_, b := t.m[templateData]
	return b
}

func (t *executor) Prepare(templateData string) error {
	tpl, err := template.New(templateData).Funcs(funcMap).Parse(templateData)
	if err != nil {
		return errors.SystemError.InvalidValidatorErrorTemplate.Wrap(err)
	}
	t.m[templateData] = tpl
	return nil
}

func (t *executor) Exec(templateData string, data interface{}) (string, error) {
	tpl, b := t.m[templateData]
	if !b {
		return "", errors.SystemError.InvalidValidatorErrorTemplate.Wrap(fmt.Errorf("unregistered template: %s", templateData))
	}
	out := bytes.NewBuffer(nil)
	if err := tpl.Execute(out, data); err != nil {
		return "", errors.SystemError.InvalidValidatorErrorTemplate.Wrap(err)
	}
	return out.String(), nil
}

func NewExecutor() Executor {
	return &executor{
		m: make(map[string]*template.Template),
	}
}
