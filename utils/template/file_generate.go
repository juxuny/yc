package template

import (
	"embed"
	"github.com/juxuny/yc/errors"
	"html/template"
	"os"
)

var funcMap = map[string]interface{}{
	"upperFirst":  toUpperFirst,
	"lowerFirst":  toLowerFirst,
	"trimPointer": trimPointer,
	"raw":         raw,
}

func RunEmbedFile(fs embed.FS, templateFileName string, outputFileName string, data interface{}) error {
	buf, err := fs.ReadFile(templateFileName)
	if err != nil {
		return errors.SystemError.FsReadTemplateDataFailed.Wrap(err)
	}
	tpl, err := template.New(templateFileName).Funcs(funcMap).Parse(string(buf))
	if err != nil {
		return errors.SystemError.TemplateSyntaxError.Wrap(err)
	}
	out, err := os.Create(outputFileName)
	if err != nil {
		return errors.SystemError.FsCreateFailed.Wrap(err)
	}
	defer func() {
		_ = out.Close()
	}()
	return tpl.Execute(out, data)
}

func AppendFromEmbedFile(fs embed.FS, templateFileName string, outputFileName string, data interface{}) error {
	buf, err := fs.ReadFile(templateFileName)
	if err != nil {
		return errors.SystemError.FsReadTemplateDataFailed.Wrap(err)
	}
	tpl, err := template.New(templateFileName).Funcs(funcMap).Parse(string(buf))
	if err != nil {
		return errors.SystemError.TemplateSyntaxError.Wrap(err)
	}
	out, err := os.OpenFile(outputFileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return errors.SystemError.FsCreateFailed.Wrap(err)
	}
	defer func() {
		_ = out.Close()
	}()
	return tpl.Execute(out, data)

}
