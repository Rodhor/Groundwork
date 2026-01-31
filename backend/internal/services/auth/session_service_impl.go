package authService

import (
	. "Groundwork/backend/internal/database"
	. "Groundwork/backend/internal/domain"

	"context"

	"github.com/google/uuid"
)

type sessionService struct {
	sessionQueries *SessionDB
}

func NewSessionService() SessionService {
	return &sessionService{
		sessionQueries: NewSessionDB(),
	}
}
func (s *sessionService) CreateSession(ctx context.Context, userID uuid.UUID) (*Session, error) {
	session, err := s.sessionQueries.CreateSession(ctx, userID)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (s *sessionService) EndSession(ctx context.Context, sessionID uuid.UUID) error {
	err := s.sessionQueries.EndSession(ctx, sessionID)
	if err != nil {
		return err
	}
	return nil
}

func (s *sessionService) GetSession(ctx context.Context, sessionID uuid.UUID) (*Session, error) {
	session, err := s.sessionQueries.GetSession(ctx, sessionID)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (s *sessionService) ResetSessionExpiration(ctx context.Context, sessionID uuid.UUID, params DurationParams) (*Session, error) {
	session, err := s.sessionQueries.ResetSessionExpiration(ctx, sessionID, params)
	if err != nil {
		return nil, err
	}
	return session, nil
}
