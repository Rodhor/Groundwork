package authCommand

import (
	. "Groundwork/backend/internal/commands"
	. "Groundwork/backend/internal/domain"
	. "Groundwork/backend/internal/services/auth"
	"context"
	"net/http"

	"github.com/google/uuid"
)

type sessionCommand struct {
	sessionService SessionService
}

var (
	creationFailed               = "failed to create session"
	sessionNotFound              = "session not found"
	sessionEndFailed             = "failed to end session"
	sessionExpirationResetFailed = "failed to reset session expiration"
)
var (
	sessionCreated         = "session created"
	sessionRetrieved       = "session retrieved"
	sessionEnded           = "session ended"
	sessionExpirationReset = "session expiration reset"
)

func NewSessionCommand(sessionService SessionService) SessionCommand {
	return &sessionCommand{
		sessionService: sessionService,
	}
}

func (c *sessionCommand) CreateSessionCommand(ctx context.Context, userID uuid.UUID) Response {
	session, err := c.sessionService.CreateSession(ctx, userID)
	if err != nil {
		return Response{
			Status:  http.StatusBadRequest,
			Message: creationFailed,
			Data:    err,
		}
	}
	return Response{
		Status:  http.StatusCreated,
		Message: sessionCreated,
		Data:    session,
	}
}

func (c *sessionCommand) GetSessionCommand(ctx context.Context, sessionID uuid.UUID) Response {
	session, err := c.sessionService.GetSession(ctx, sessionID)
	if err != nil {
		return Response{
			Status:  http.StatusNotFound,
			Message: sessionNotFound,
			Data:    err,
		}
	}
	return Response{
		Status:  http.StatusOK,
		Message: sessionRetrieved,
		Data:    session,
	}
}

func (c *sessionCommand) EndSessionCommand(ctx context.Context, sessionID uuid.UUID) Response {
	err := c.sessionService.EndSession(ctx, sessionID)
	if err != nil {
		return Response{
			Status:  http.StatusInternalServerError,
			Message: sessionEndFailed,
			Data:    err,
		}
	}
	return Response{
		Status:  http.StatusNoContent,
		Message: sessionEnded,
		Data:    nil,
	}
}

func (c *sessionCommand) ResetSessionExpirationCommand(ctx context.Context, sessionID uuid.UUID, params DurationParams) Response {
	session, err := c.sessionService.ResetSessionExpiration(ctx, sessionID, params)
	if err != nil {
		return Response{
			Status:  http.StatusInternalServerError,
			Message: sessionExpirationResetFailed,
			Data:    err,
		}
	}
	return Response{
		Status:  http.StatusNoContent,
		Message: sessionExpirationReset,
		Data:    session,
	}
}
