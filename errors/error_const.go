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
	DuplicatedConfigName          Error `code:"-10017" msg:"duplicated config name"`
	DatabaseConnectError          Error `code:"-10018" msg:"database connect error"`
	DatabaseConfigNotFound        Error `code:"-10019" msg:"database config not found"`
	DatabaseConnectionIndexError  Error `code:"-10020" msg:"database connection index error"`
	DatabaseQueryError            Error `code:"-10021" msg:"database query error"`
	DatabaseExecError             Error `code:"-10022" msg:"database exec error"`
	DatabaseColumnError           Error `code:"-10023" msg:"database column error"`
	DatabaseColumnTypeError       Error `code:"-10024" msg:"database column type error"`
	DatabaseScanError             Error `code:"-10025" msg:"scan error"`
	ReflectNoFieldError           Error `code:"-10026" msg:"field not exists"`
	DatabaseNoData                Error `code:"-10027" msg:"no data"`
	NotSupportedDataType          Error `code:"-10028" msg:"not supported data type"`
	NotPointer                    Error `code:"-10029" msg:"not pointer"`
	InvalidGoModule               Error `code:"-10030" msg:"invalid go module"`
	NotFoundModuleFile            Error `code:"-10031" msg:"not found go.mod file"`
	InternalError                 Error `code:"-10032" msg:"server error"`
	RpcCallLevelNotAllow          Error `code:"-10033" msg:"rpc level not allow"`
	RpcCallMetaEmpty              Error `code:"-10034" msg:"rpc call metadata empty"`
	InvalidRpcCallerLevel         Error `code:"-10035" msg:"invalid rpc caller level"`
	NotFoundRpcCallerLevel        Error `code:"-10036" msg:"not found rpc caller level"`
	NotFoundRpcToken              Error `code:"-10037" msg:"no token"`
	InvalidToken                  Error `code:"-10038" msg:"invalid token"`
	NoFields                      Error `code:"-10039" msg:"no fields"`
}{}

func init() {
	if err := InitErrorStruct(&SystemError); err != nil {
		panic(err)
	}
}
