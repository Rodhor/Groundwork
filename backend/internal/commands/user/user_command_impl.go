package internalUserCommand

import (
	. "Groundwork/backend/internal/commands"
	. "Groundwork/backend/internal/domain"
	. "Groundwork/backend/internal/services/user"
	"context"
	"errors"
	"net/http"
)

// Error messages
var (
	CreationFailed = errors.New("user creation failed").Error()
	NotFound       = errors.New("user not found").Error()
	DeletionFailed = errors.New("user deletion failed").Error()
	UpdateFailed   = errors.New("user update failed").Error()
)

// Success messages
var (
	Created = "User was created succesfully"
	Updated = "User was updated succesfully"
	Deleted = "User was deleted succesfully"
	Found   = "User was found succesfully"
)

type userCommand struct {
	userService UserService
}

func NewUserCommand(userService UserService) UserCommand {
	return &userCommand{
		userService: userService,
	}
}

func (u *userCommand) AddNewUser_command(ctx context.Context, user *User) Response {
	newUser, errs := u.userService.AddNewUser_service(ctx, user)
	if errs != nil {
		return Response{
			Status:  http.StatusBadRequest,
			Message: CreationFailed,
			Data:    errs,
		}
	}
	return Response{
		Status:  http.StatusCreated,
		Message: Created,
		Data:    newUser,
	}
}

func (u *userCommand) GetUserByID_command(ctx context.Context, id int64) Response {
	user, err := u.userService.GetUserByID_service(ctx, id)
	if err != nil {
		return Response{
			Status:  http.StatusNotFound,
			Message: NotFound,
			Data:    err,
		}
	}
	return Response{
		Status:  http.StatusOK,
		Message: Found,
		Data:    user,
	}
}

func (u *userCommand) GetUserByUsername_command(ctx context.Context, username string) Response {
	user, err := u.userService.GetUserByUsername_service(ctx, username)
	if err != nil {
		return Response{
			Status:  http.StatusNotFound,
			Message: NotFound,
			Data:    err,
		}
	}
	return Response{
		Status:  http.StatusOK,
		Message: Found,
		Data:    user,
	}
}

func (u *userCommand) UpdateUser_command(ctx context.Context, user *User) Response {
	errs := u.userService.UpdateUser_service(ctx, user)
	if errs != nil {
		return Response{
			Status:  http.StatusBadRequest,
			Message: UpdateFailed,
			Data:    errs,
		}
	}
	return Response{
		Status:  http.StatusOK,
		Message: Updated,
		Data:    user,
	}
}

func (u *userCommand) DeleteUser_command(ctx context.Context, id int64) Response {
	err := u.userService.DeleteUser_service(ctx, id)
	if err != nil {
		return Response{
			Status:  http.StatusBadRequest,
			Message: DeletionFailed,
			Data:    err,
		}
	}
	return Response{
		Status:  http.StatusOK,
		Message: Deleted,
		Data:    nil,
	}
}
