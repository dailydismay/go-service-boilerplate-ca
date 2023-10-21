package pgsql

import (
	"authsvc/internal/domain/user"
	"time"
)

type userRow struct {
	ID               string     `db:"id"`
	Username         string     `db:"username"`
	Password         string     `db:"password"`
	Created          *time.Time `db:"created_at"`
	LastLoginAttempt *time.Time `db:"last_login_attempt"`
}

func (r *userRow) ToDomain() *user.User {
	return user.FromData(user.ID(r.ID), r.Username, r.Password, r.Created, r.LastLoginAttempt)
}
