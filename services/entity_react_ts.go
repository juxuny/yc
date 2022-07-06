package services

type EnumField struct {
	FieldName string
	Value     string
	Desc      string
}

type EnumEntity struct {
	ServiceEntity
	EnumName          string
	Fields            []EnumField
	Desc              string
	ValueSet          string
	CSharperNamespace string
}

type ReactTsMessageField struct {
	Name     string
	Desc     string
	Required bool
	Type     string
}

type ReactTsMessage struct {
	Desc   string
	Name   string
	Fields []ReactTsMessageField
}

type ReactTsTypingEntity struct {
	ServiceEntity
	Desc     string
	Enums    []EnumEntity
	Messages []ReactTsMessage
}

type ReactTsMethod struct {
	ServiceEntity
	MethodName string
	Request    string
	Response   string
	Desc       string
	Api        string
}

type ReactTsApiEntity struct {
	ServiceEntity
	Methods []ReactTsMethod
}
