package UserService

import (
	. "Groundwork/backend/internal/database"
	. "Groundwork/backend/internal/domain"
	"context"
	"log/slog"

	"github.com/google/uuid"
)

type userService struct {
	queries *UserDB
}

func NewUserService() UserService {
	return &userService{
		NewUserDB(),
	}
}

func (s *userService) AddNewUser_service(ctx context.Context, user *User) (*User, map[string]string) {
	// check if user is valid
	if errors := user.Validate(); len(errors) > 0 {
		return nil, errors
	}
	// set user initials to first two letters of name
	user.SetUserInitials()

	slog.Info("Adding new user", "user", user)
	createdUser, err := s.queries.AddNewUser(ctx, user)
	if err != nil {
		return nil, map[string]string{"error": err.Error()}

	}
	return createdUser, nil
}

func (s *userService) GetUserByID_service(ctx context.Context, id uuid.UUID) (*User, error) {
	slog.Info("Getting user by ID", "id", id)
	user, err := s.queries.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) GetUserByUsername_service(ctx context.Context, username string) (*User, error) {
	slog.Info("Getting user by username", "username", username)
	user, err := s.queries.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) UpdateUser_service(ctx context.Context, user *User) map[string]string {
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

func (s *userService) DeleteUser_service(ctx context.Context, id uuid.UUID) error {
	slog.Info("Deleting user", "id", id)
	return s.queries.DeleteUser(ctx, id)
}
