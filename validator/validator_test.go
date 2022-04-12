package validator

import (
	"testing"
)

func TestValidator(t *testing.T) {
	type Level int
	type TestStruct struct {
		Level   Level
		Content string
		List    []int
	}
	v := &TestStruct{
		Level:   5,
		Content: "100",
		List:    []int{1},
	}
	action := Action{
		ValidatorFormulas: []string{
			"min", "max",
		},
		ErrorTemplate: "{{if lt .Level 5r }}minimum of level is 5, current value is: {{.Level}}{{else}}maximum of level is 10{{end}}",
	}
	if err := RegisterTemplate(action.ErrorTemplate); err != nil {
		t.Fatal(err)
	}
	if err := Run(v.Level, "5", action, v); err != nil {
		t.Fatal(err)
	}
}
