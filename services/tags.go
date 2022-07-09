package services

type ProtoTag string

const (
	ProtoTagInternal      = ProtoTag("@internal")
	ProtoTagGroup         = ProtoTag("@group")
	ProtoTagLevel         = ProtoTag("@level")
	ProtoTagValidator     = ProtoTag("@v")
	ProtoTagDesc          = ProtoTag("@desc")
	ProtoTagName          = ProtoTag("@name")
	ProtoTagRef           = ProtoTag("@ref")
	ProtoTagMsg           = ProtoTag("@msg")
	ProtoTagIndex         = ProtoTag("@index")
	ProtoTagOrm           = ProtoTag("@orm")
	ProtoTagPrimary       = ProtoTag("@primary")
	ProtoTagIgnoreProto   = ProtoTag("@ignore-proto")
	ProtoTagUnique        = ProtoTag("@unique")
	ProtoTagAuth          = ProtoTag("@auth")
	ProtoTagIgnoreAuth    = ProtoTag("@ignore-auth")
	ProtoTagAutoIncrement = ProtoTag("@auto-increment")
	ProtoTagVersion       = ProtoTag("@version")
)

func (t ProtoTag) String() string {
	return string(t)
}
