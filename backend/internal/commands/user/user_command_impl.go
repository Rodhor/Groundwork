package internalUserCommand

import (
	. "Groundwork/backend/internal/commands"
	"Groundwork/backend/internal/domain"
	. "Groundwork/backend/internal/services/user"
	"context"
	"net/http"
)

type userCommand struct {
	userService UserService
}

func NewUserCommand(userService UserService) UserCommand {
	return &userCommand{
		userService: userService,
	}
}

func (u *userCommand) AddNewUser_command(ctx context.Context, user *domain.User) Response {
	newUser, errs := u.userService.AddNewUser_service(ctx, user)
	if errs != nil {
		return Response{
			Status:  http.StatusBadRequest,
			Message: "An error occured while creating the user.",
			Data:    errs,
		}
	}
	return Response{
		Status:  http.StatusCreated,
		Message: "User was created succesfully",
		Data:    newUser,
	}
}

func (u *userCommand) GetUserByID_command(ctx context.Context, id int64) Response {
	user, err := u.userService.GetUserByID_service(ctx, id)
	if err != nil {
		return Response{
			Status:  http.StatusNotFound,
			Message: "User not found",
			Data:    err,
		}
	}
	return Response{
		Status:  http.StatusOK,
		Message: "User found",
		Data:    user,
	}
}

func (u *userCommand) UpdateUser_command(ctx context.Context, user *domain.User) Response {
	errs := u.userService.UpdateUser_service(ctx, user)
	if errs != nil {
		return Response{
			Status:  http.StatusBadRequest,
			Message: "An error occured while updating the user.",
			Data:    errs,
		}
	}
	return Response{
		Status:  http.StatusOK,
		Message: "User updated succesfully",
		Data:    user,
	}
}

func (u *userCommand) DeleteUser_command(ctx context.Context, id int64) Response {
	err := u.userService.DeleteUser_service(ctx, id)
	if err != nil {
		return Response{
			Status:  http.StatusBadRequest,
			Message: "An error occured while deleting the user.",
			Data:    err,
		}
	}
	return Response{
		Status:  http.StatusOK,
		Message: "User deleted succesfully",
		Data:    nil,
	}
}
