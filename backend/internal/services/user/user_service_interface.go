package internalUserService

import (
	. "Groundwork/backend/internal/domain"
	"context"
)

type UserService interface {
	AddNewUser_service(ctx context.Context, user *User) (*User, map[string]string)
	GetUserByID_service(ctx context.Context, id int64) (*User, error)
	GetUserByUsername_service(ctx context.Context, username string) (*User, error)
	UpdateUser_service(ctx context.Context, user *User) map[string]string
	DeleteUser_service(ctx context.Context, id int64) error
}
