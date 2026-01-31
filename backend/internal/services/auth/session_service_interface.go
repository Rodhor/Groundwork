package authService

import (
	. "Groundwork/backend/internal/domain"
	"context"

	"github.com/google/uuid"
)

type SessionService interface {
	CreateSession(ctx context.Context, userID uuid.UUID) (*Session, error)
	GetSession(ctx context.Context, sessionID uuid.UUID) (*Session, error)
	EndSession(ctx context.Context, sessionID uuid.UUID) error
	ResetSessionExpiration(ctx context.Context, sessionID uuid.UUID, params DurationParams) (*Session, error)
}
