package services

type ProtoTag string

const (
	ProtoTagInternal  = ProtoTag("@internal")
	ProtoTagGroup     = ProtoTag("@group")
	ProtoTagLevel     = ProtoTag("@level")
	ProtoTagValidator = ProtoTag("@v")
	ProtoTagDesc      = ProtoTag("@desc")
	ProtoTagName      = ProtoTag("@name")
	ProtoTagRef       = ProtoTag("@ref")
	ProtoTagMsg       = ProtoTag("@msg")
	ProtoTagIndex     = ProtoTag("@index")
	ProtoTagOrm       = ProtoTag("@orm")
)

func (t ProtoTag) String() string {
	return string(t)
}
