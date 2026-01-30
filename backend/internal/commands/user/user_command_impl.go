package internalUserCommand

import (
	. "Groundwork/backend/internal/commands"
	"Groundwork/backend/internal/domain"
	internalUserService "Groundwork/backend/internal/services/user"
	"context"
)

type userCommand struct {
	userService internalUserService.UserService
}

func NewUserCommand(userService internalUserService.UserService) UserCommand {
	return &userCommand{
		userService: userService,
	}
}
func (u *userCommand) AddNewUser_command(ctx context.Context, user *domain.User) Response {
	panic("unimplemented")
}

func (u *userCommand) DeleteUser_command(ctx context.Context, id int64) Response {
	panic("unimplemented")
}

func (u *userCommand) GetUserByID_command(ctx context.Context, id int64) Response {
	panic("unimplemented")
}

func (u *userCommand) UpdateUser_command(ctx context.Context, user *domain.User) Response {
	panic("unimplemented")
}
