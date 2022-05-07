package dt

type Role string

const (
	RoleSuper = Role("super")
	RoleAdmin = Role("admin")
)

func (t Role) String() string {
	return string(t)
}
