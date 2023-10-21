package user

import (
	"time"

	"github.com/google/uuid"
)

type ID string

type User struct {
	ID

	Username, Password string

	CreatedAt, LastLoginAttemp *time.Time
}

func New(username, password string) *User {
	now := time.Now()
	return &User{
		ID:        ID(uuid.New().String()),
		Username:  username,
		Password:  password,
		CreatedAt: &now,
	}
}

func FromData(id ID, username, password string, created, lastLogin *time.Time) *User {
	return &User{
		id,
		username,
		password,
		created,
		lastLogin,
	}
}
