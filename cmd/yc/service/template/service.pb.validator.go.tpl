package {{.PackageName}}

import (
	"github.com/juxuny/yc/validator"
)

const (
{{range $item := .Messages}}{{range $f := .Fields}}	ValidatorTemplate{{$item.Name}}{{$f.Name}} = "{{$f.Error}}"
{{end}}{{end}})

var templateList = []string{
{{range $item := .Messages}}{{range $f := .Fields}}	ValidatorTemplate{{$item.Name}}{{$f.Name}},
{{end}}{{end}}}

func init() {
	if err := validator.RegisterTemplate(templateList...); err != nil {
		panic(err)
	}
}
{{range $msg := .Messages}}
func (m *{{$msg.Name}}) Validate() error {
{{range $f := .Fields}}{{range $formula := $f.Formulas}}	if err := validator.Run(m.{{$f.Name}}, validator.CreateAction("{{$formula.Pattern}}", "{{$formula.RefValue}}", ValidatorTemplate{{$msg.Name}}{{$f.Name}}), m, "{{$f.ParamName}}"); err != nil {
		return err
	}
{{end}}{{end}}	return nil
}
{{end}}
