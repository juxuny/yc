package errors

var SystemError = struct {
	FsIsNotDir                   Error `code:"-1000" msg:"is not a directory"`
	FsReadTemplateDataFailed     Error `code:"-1001" msg:"read template file error"`
	FsCreateFailed               Error `code:"-1002" msg:"create file error"`
	TemplateSyntaxError          Error `code:"-1003" msg:"template syntax error"`
	LogDirEmpty                  Error `code:"-1004" msg:"log directory is empty"`
	NotFound                     Error `code:"-1005" msg:"not found"`
	NotSupportedMethod           Error `code:"-1006" msg:"not supported method"`
	InvalidNumberOfParams        Error `code:"-1007" msg:"invalid number of params"`
	InvalidNumberOfReplyEntities Error `code:"-1008" msg:"invalid number of reply entities"`
	InvalidInputDataObject       Error `code:"-1009" msg:"invalid input data object"`
	InvalidFormData              Error `code:"-1010" msg:"invalid form data"`
	InvalidJsonData              Error `code:"-1011" msg:"invalid json data"`
}{}

func init() {
	if err := InitErrorStruct(&SystemError); err != nil {
		panic(err)
	}
}
