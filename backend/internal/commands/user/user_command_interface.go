package internalUserCommand

import (
	commandUtils "Groundwork/backend/internal/commands"
	"Groundwork/backend/internal/domain"
	"context"
)

type UserCommand interface {
	AddNewUser_command(ctx context.Context, user *domain.User) commandUtils.Response
	GetUserByID_command(ctx context.Context, id int64) commandUtils.Response
	UpdateUser_command(ctx context.Context, user *domain.User) commandUtils.Response
	DeleteUser_command(ctx context.Context, id int64) commandUtils.Response
}
