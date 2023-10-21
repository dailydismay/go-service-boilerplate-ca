package session

import (
	"authsvc/internal/domain/user"
	"time"
)

type Session struct {
	ID string

	UserID    user.ID
	ExpiresAt time.Time
}

func New(id string, userID user.ID, expiresAt time.Time) *Session {
	return &Session{id, userID, expiresAt}
}
