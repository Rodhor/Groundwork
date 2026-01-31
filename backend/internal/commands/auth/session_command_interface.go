package authCommand

import (
	. "Groundwork/backend/internal/commands"
	. "Groundwork/backend/internal/domain"
	"context"

	"github.com/google/uuid"
)

type SessionCommand interface {
	CreateSessionCommand(ctx context.Context, userID uuid.UUID) Response
	GetSessionCommand(ctx context.Context, sessionID uuid.UUID) Response
	EndSessionCommand(ctx context.Context, sessionID uuid.UUID) Response
	ResetSessionExpirationCommand(ctx context.Context, sessionID uuid.UUID, params DurationParams) Response
}
