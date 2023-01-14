package template

import (
	"embed"
	"github.com/juxuny/yc/errors"
	"html/template"
	"io/ioutil"
	"os"
)

var funcMap = map[string]interface{}{
	"upperFirst":  toUpperFirst,
	"lowerFirst":  toLowerFirst,
	"trimPointer": trimPointer,
	"raw":         raw,
	"num":         castNumber,
	"lower":       lower,
	"upper":       upper,
	"camelcase":   camelcaseString,
	"underline":   underlineString,
	"inc":         increase,
	"dec":         decrease,
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

// SaveEmbedFileAs save embed file to a specified path
func SaveEmbedFileAs(fs embed.FS, fileName string, outputFileName string) error {
	buf, err := fs.ReadFile(fileName)
	if err != nil {
		return errors.SystemError.FsReadTemplateDataFailed.Wrap(err)
	}
	err = ioutil.WriteFile(outputFileName, buf, 0666)
	if err != nil {
		return errors.SystemError.FsWriteError.Wrap(err)
	}
	return nil
}
