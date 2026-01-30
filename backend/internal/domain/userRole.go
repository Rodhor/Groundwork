package domain

type UserRole string

const (
	Admin    UserRole = "administrator"
	Operator UserRole = "operator"
	Analyst  UserRole = "analyst"
)

func (r UserRole) IsValid() bool {
	switch r {
	case Admin, Operator, Analyst:
		return true
	}
	return false
}
