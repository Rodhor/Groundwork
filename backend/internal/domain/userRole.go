package domain

type UserRole int

const (
	Admin UserRole = iota
	Operator
	Analyst
)

func (r UserRole) IsValid() bool {
	switch r {
	case Admin, Operator, Analyst:
		return true
	}
	return false
}

func (r UserRole) String() string {
	switch r {
	case Admin:
		return "Admin"
	case Operator:
		return "Operator"
	case Analyst:
		return "Analyst"
	default:
		return "Unknown"
	}
}
