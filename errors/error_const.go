package errors

var SystemError = struct {
	FsIsNotDir                    Error `code:"-10000" msg:"is not a directory"`
	FsReadTemplateDataFailed      Error `code:"-10001" msg:"read template file error"`
	FsCreateFailed                Error `code:"-10002" msg:"create file error"`
	TemplateSyntaxError           Error `code:"-10003" msg:"template syntax error"`
	LogDirEmpty                   Error `code:"-10004" msg:"log directory is empty"`
	NotFound                      Error `code:"-10005" msg:"not found"`
	NotSupportedMethod            Error `code:"-10006" msg:"not supported method"`
	InvalidNumberOfParams         Error `code:"-10007" msg:"invalid number of params"`
	InvalidNumberOfReplyEntities  Error `code:"-10008" msg:"invalid number of reply entities"`
	InvalidInputDataObject        Error `code:"-10009" msg:"invalid input data object"`
	InvalidFormData               Error `code:"-10010" msg:"invalid form data"`
	InvalidJsonData               Error `code:"-10011" msg:"invalid json data"`
	InvalidValidatorFormula       Error `code:"-10012" msg:"invalid validator formula"`
	InvalidValidatorErrorTemplate Error `code:"-10013" msg:"invalid error template"`
	InvalidParams                 Error `code:"-10014" msg:"invalid params"`
	InvalidDataType               Error `code:"-10015" msg:"invalid data type"`
	InvalidRefValueDefinition     Error `code:"-10016" msg:"invalid ref value definition"`
}{}

func init() {
	if err := InitErrorStruct(&SystemError); err != nil {
		panic(err)
	}
}
