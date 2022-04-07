package errors

var SystemError = struct {
	FsIsNotDir               Error `code:"-1000" msg:"is not a directory"`
	FsReadTemplateDataFailed Error `code:"-1001" msg:"read template file error"`
	FsCreateFailed           Error `code:"-1002" msg:"create file error"`
	TemplateSyntaxError      Error `code:"-1003" msg:"template syntax error"`
}{}

func init() {
	if err := InitErrorStruct(&SystemError); err != nil {
		panic(err)
	}
}
