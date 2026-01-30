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

func (s *userService) AddNewUser_service(ctx context.Context, user *domain.User) (*domain.User, map[string]string) {
	// check if user is valid
	if errors := user.Validate(); len(errors) > 0 {
		return nil, errors
	}
	slog.Info("Adding new user", "user", user)
	createdUser, err := s.queries.AddNewUser(ctx, user)

	if err != nil {
		return nil, map[string]string{"error": err.Error()}

	}
	return createdUser, nil
}

func (s *userService) GetUserByID_service(ctx context.Context, id int64) (*domain.User, error) {
	slog.Info("Getting user by ID", "id", id)
	return s.queries.GetUserByID(ctx, id)
}

func (s *userService) GetUserByUsername_service(ctx context.Context, username string) (*domain.User, error) {
	slog.Info("Getting user by username", "username", username)
	return s.queries.GetUserByUsername(ctx, username)
}

func (s *userService) UpdateUser_service(ctx context.Context, user *domain.User) map[string]string {
	// check if user is valid
	if errs := user.Validate(); len(errs) > 0 {
		return errs
	}
	slog.Info("Updating user", "user", user)
	err := s.queries.UpdateUser(ctx, user)
	if err != nil {
		return map[string]string{"error": err.Error()}
	}
	return nil
}

func (s *userService) DeleteUser_service(ctx context.Context, id int64) error {
	slog.Info("Deleting user", "id", id)
	return s.queries.DeleteUser(ctx, id)
}
