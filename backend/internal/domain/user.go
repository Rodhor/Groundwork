package domain

import (
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

func (u *User) Validate() map[string]string {

	// Map to hold validation errors
	errs := make(map[string]string)

	if u.Email == "" {
		errs["email"] = "email cannot be empty"
	} else if _, err := mail.ParseAddress(u.Email); err != nil {
		errs["email"] = "invalid email format"
	}

	// Validate that username is not empty
	if u.Username == "" {
		errs["username"] = "username cannot be empty"
	}

	// Validate that name is not empty
	if u.Name == "" {
		errs["name"] = "name cannot be empty"
	}

	// Validate that password is not empty
	if u.PasswordHash == "" {
		errs["password"] = "password cannot be empty"
	}

	// Validate that role is valid
	if !u.Role.IsValid() {
		errs["role"] = "invalid user role"
	}

	// Return validation errors even if empty (no errors)
	return errs

}
