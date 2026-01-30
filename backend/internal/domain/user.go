package domain

import (
	"errors"
	"net/mail"
	"time"
)

type User struct {
	ID           int64
	Name         string
	Username     string
	PasswordHash string
	Email        string
	Role         UserRole
	LastLogin    time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
	CreatedBy    int64
	UpdatedBy    int64
}

func (u *User) Validate() map[string]error {

	// Map to hold validation errors
	errs := make(map[string]error)

	// Validate that email is not empty
	if u.Email == "" {
		errs["email"] = errors.New("email cannot be empty")
	}

	// Validate email format
	_, err := mail.ParseAddress(u.Email)

	if err != nil {
		errs["email"] = errors.New("invalid email format")
	}

	// Validate that username is not empty
	if u.Username == "" {
		errs["username"] = errors.New("username cannot be empty")
	}

	// Validate that name is not empty
	if u.Name == "" {
		errs["name"] = errors.New("name cannot be empty")
	}

	// Validate that password is not empty
	if u.PasswordHash == "" {
		errs["password"] = errors.New("password cannot be empty")
	}

	// Validate that role is valid
	if !u.Role.IsValid() {
		errs["role"] = errors.New("invalid user role")
	}

	// Return validation errors even if empty (no errors)
	return errs

}
