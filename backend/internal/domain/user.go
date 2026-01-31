package domain

import (
	"errors"
	"net/mail"
	"time"
)

type User struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	Username     string    `json:"username"`
	PasswordHash string    `json:"password_hash"`
	Email        string    `json:"email"`
	Role         UserRole  `json:"role"`
	LastLogin    time.Time `json:"last_login"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	CreatedBy    int64     `json:"created_by"`
	UpdatedBy    int64     `json:"updated_by"`
}

var (
	InvalidEmailFormat    = errors.New("invalid email format").Error()
	InvalidRole           = errors.New("invalid role").Error()
	InvalidNameEmpty      = errors.New("name cannot be empty").Error()
	InvalidUsernameEmpty  = errors.New("username cannot be empty").Error()
	InvalidUsernameFormat = errors.New("invalid username format").Error()
	InvalidEmailEmpty     = errors.New("email cannot be empty").Error()
	InvalidPasswordEmpty  = errors.New("password cannot be empty").Error()
)

func (u *User) Validate() map[string]string {

	// Map to hold validation errors
	errs := make(map[string]string)

	if u.Email == "" {
		errs["email"] = InvalidEmailEmpty
	} else if _, err := mail.ParseAddress(u.Email); err != nil {
		errs["email"] = InvalidEmailFormat
	}

	// Validate that username is not empty
	if u.Username == "" {
		errs["username"] = InvalidUsernameEmpty
	}

	// Validate that name is not empty
	if u.Name == "" {
		errs["name"] = InvalidNameEmpty
	}

	// Validate that password is not empty
	if u.PasswordHash == "" {
		errs["password"] = InvalidPasswordEmpty
	}

	// Validate that role is valid
	if !u.Role.IsValid() {
		errs["role"] = InvalidRole
	}

	// Return validation errors even if empty (no errors)
	return errs

}
