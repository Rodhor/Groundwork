package internalUserService

import (
	"Groundwork/backend/internal/domain"
	"context"
)

type UserService interface {
	AddNewUser_service(ctx context.Context, user *domain.User) (*domain.User, map[string]string)
	GetUserByID_service(ctx context.Context, id int64) (*domain.User, error)
	UpdateUser_service(ctx context.Context, user *domain.User) map[string]string
	DeleteUser_service(ctx context.Context, id int64) error
}
