package domain

import (
	"time"

	"github.com/google/uuid"
)

type Format string

var (
	Hour        Format = "hour"
	Minute      Format = "minute"
	Second      Format = "second"
	MilliSecond Format = "millisecond"
)

type Session struct {
	ID        uuid.UUID `json:"session_id"`
	UserID    uuid.UUID `json:"user_id"`
	ExpiresAt time.Time `json:"expires_at"`
	StartedAt time.Time `json:"started_at"`
	EndedAt   time.Time `json:"ended_at"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DurationParams struct {
	Format   *Format `json:"format"`
	Duration *int64  `json:"duration"`
}

func (f Format) isValid() bool {
	switch f {
	case Hour, Minute, Second, MilliSecond:
		return true
	default:
		return false
	}
}

func (s *Session) EndSession() {
	s.EndedAt = time.Now()
	s.IsActive = false
}

func (s *Session) IsExpired() bool {
	return s.ExpiresAt.Before(time.Now())
}

func (s *Session) ResetExpiration(params DurationParams) {
	format := Hour
	duration := int64(24)

	if params.Format != nil && params.Format.isValid() {
		format = *params.Format
	}
	if params.Duration != nil && *params.Duration > 0 {
		duration = *params.Duration
	}

	switch format {
	case Hour:
		s.ExpiresAt = time.Now().Add(time.Hour * time.Duration(duration))
	case Minute:
		s.ExpiresAt = time.Now().Add(time.Minute * time.Duration(duration))
	case Second:
		s.ExpiresAt = time.Now().Add(time.Second * time.Duration(duration))
	case MilliSecond:
		s.ExpiresAt = time.Now().Add(time.Millisecond * time.Duration(duration))
	default:
		s.ExpiresAt = time.Now().Add(time.Hour * 24)
	}
}
