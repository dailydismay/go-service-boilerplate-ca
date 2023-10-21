package token

import (
	"authsvc/internal/domain/user"
	"time"

	"github.com/google/uuid"
)

type Claims struct {
	SessID    string
	UserID    user.ID
	ExpiresAt time.Time
}

func NewClaimsFactory(sessID string, userID user.ID) func(expiresAt time.Time) *Claims {
	return func(expiresAt time.Time) *Claims {
		return &Claims{sessID, userID, expiresAt}
	}
}

type Token struct {
	ID, Value string
	ExpiresAt time.Time
}

func New(expiresAt time.Time) *Token {
	return &Token{
		ID:        uuid.New().String(),
		ExpiresAt: expiresAt,
	}
}
