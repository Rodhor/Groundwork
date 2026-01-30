package internalUserService

import (
	"Groundwork/backend/internal/database"
	"Groundwork/backend/internal/domain"
	"context"
	"log/slog"
)

type userService struct {
	queries *database.UserDB
}

func NewUserService() UserService {
	return &userService{
		database.NewUserDB(),
	}
}

func (s *userService) AddNewUser_service(ctx context.Context, user *domain.User) (*domain.User, error) {
	slog.Info("Adding new user", "user", user)
	return s.queries.AddNewUser(ctx, user)
}

func (s *userService) GetUserByID_service(ctx context.Context, id int64) (*domain.User, error) {
	slog.Info("Getting user by ID", "id", id)
	return s.queries.GetUserByID(ctx, id)
}

func (s *userService) UpdateUser_service(ctx context.Context, user *domain.User) error {
	slog.Info("Updating user", "user", user)
	return s.queries.UpdateUser(ctx, user)
}

func (s *userService) DeleteUser_service(ctx context.Context, id int64) error {
	slog.Info("Deleting user", "id", id)
	return s.queries.DeleteUser(ctx, id)
}
