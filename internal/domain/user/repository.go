package user

import (
	"context"
	"time"
)

type Repository interface {
	Save(context.Context, *User) error
	FindByUsername(context.Context, string) (*User, error)
	UpdateLastLoginAttempt(context.Context, ID, time.Time) error
}
