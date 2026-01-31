package UserCommand

import (
	. "Groundwork/backend/internal/commands"
	. "Groundwork/backend/internal/domain"
	"context"

	"github.com/google/uuid"
)

type UserCommand interface {
	AddNewUser_command(ctx context.Context, user *User) Response
	GetUserByID_command(ctx context.Context, id uuid.UUID) Response
	GetUserByUsername_command(ctx context.Context, username string) Response
	UpdateUser_command(ctx context.Context, user *User) Response
	DeleteUser_command(ctx context.Context, id uuid.UUID) Response
}
