package internalUserCommand

import (
	. "Groundwork/backend/internal/commands"
	"Groundwork/backend/internal/domain"
	"context"
)

type UserCommand interface {
	AddNewUser_command(ctx context.Context, user *domain.User) Response
	GetUserByID_command(ctx context.Context, id int64) Response
	GetUserByUsername_command(ctx context.Context, username string) Response
	UpdateUser_command(ctx context.Context, user *domain.User) Response
	DeleteUser_command(ctx context.Context, id int64) Response
}
