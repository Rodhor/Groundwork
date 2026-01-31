package database

import (
	"context"
	"errors"
	"sync"
	"time"

	. "Groundwork/backend/internal/domain"

	"github.com/google/uuid"
)

type SessionDB struct {
	mu       sync.RWMutex
	Sessions map[uuid.UUID]*Session
}

// Error messages
var (
	ErrSessionNotFound = errors.New("session not found")
	ErrSessionExists   = errors.New("session already exists")
	ErrSessionInvalid  = errors.New("session invalid")
	ErrSessionExpired  = errors.New("session expired")
)

func NewSessionDB() *SessionDB {
	return &SessionDB{
		Sessions: make(map[uuid.UUID]*Session),
	}
}

func (sdb *SessionDB) checkIfSessionIsValid(sessionID uuid.UUID) error {
	session, ok := sdb.Sessions[sessionID]
	if !ok {
		return ErrSessionNotFound
	}
	if session.IsActive != true {
		return ErrSessionInvalid
	}
	if session.IsExpired() {
		return ErrSessionExpired
	}
	return nil
}

func (sdb *SessionDB) CreateSession(ctx context.Context, userID uuid.UUID) (*Session, error) {
	sessionID := uuid.New()
	session := &Session{
		ID:        sessionID,
		UserID:    userID,
		ExpiresAt: time.Now().Add(time.Hour * 24),
		StartedAt: time.Now(),
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	sdb.mu.Lock()
	defer sdb.mu.Unlock()
	if _, ok := sdb.Sessions[sessionID]; ok {
		return nil, ErrSessionExists
	}
	sdb.Sessions[sessionID] = session
	return session, nil
}

func (sdb *SessionDB) GetSession(ctx context.Context, sessionID uuid.UUID) (*Session, error) {
	sdb.mu.RLock()
	defer sdb.mu.RUnlock()
	if err := sdb.checkIfSessionIsValid(sessionID); err != nil {
		return nil, err
	}
	session, _ := sdb.Sessions[sessionID]
	return session, nil
}

func (sdb *SessionDB) EndSession(ctx context.Context, sessionID uuid.UUID) error {
	sdb.mu.Lock()
	defer sdb.mu.Unlock()
	if err := sdb.checkIfSessionIsValid(sessionID); err != nil {
		return err
	}
	session, _ := sdb.Sessions[sessionID]
	session.EndSession()
	return nil
}

func (sdb *SessionDB) ResetSessionExpiration(ctx context.Context, sessionID uuid.UUID, params DurationParams) (*Session, error) {
	sdb.mu.Lock()
	defer sdb.mu.Unlock()
	if err := sdb.checkIfSessionIsValid(sessionID); err != nil {
		return nil, err
	}
	session, _ := sdb.Sessions[sessionID]
	session.ResetExpiration(params)
	return session, nil
}
