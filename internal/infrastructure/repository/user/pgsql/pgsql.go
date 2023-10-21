package pgsql

import (
	"authsvc/internal/domain/user"
	"authsvc/internal/infrastructure/tx/pgsqltx"
	"context"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type implementation struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) user.Repository {
	return &implementation{db}
}

/*
I simplified this part but I suggest placing an extra package QUERIES WITH:
- squirrel query builder
- database schema constants (such as tables)
*/
func (i *implementation) Save(ctx context.Context, u *user.User) error {
	const insertUserQuery = `
		INSERT INTO users (id, username, password) VALUES($1, $2, $3);
	`

	q := pgsqltx.QuerierFromCtx(ctx, i.db)
	if _, err := q.ExecContext(ctx, insertUserQuery, u.ID, u.Username, u.Password); err != nil {
		return errors.Wrap(err, "failed to insert new user recotd")
	}

	return nil
}

func (i *implementation) FindByUsername(ctx context.Context, username string) (*user.User, error) {
	const selectUserByUsernameQuery = `
		SELECT * FROM users WHERE username = $1;
	`

	q := pgsqltx.QuerierFromCtx(ctx, i.db)
	var row userRow
	err := q.GetContext(ctx, &row, selectUserByUsernameQuery, username)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.Wrap(user.ErrUserNotFound, "user not found in pg repo")
		default:
			return nil, errors.Wrap(err, "unexpected query error")
		}
	}

	return row.ToDomain(), nil
}

func (i *implementation) UpdateLastLoginAttempt(ctx context.Context, id user.ID, updatedDate time.Time) error {
	const updateUserLastLoginAttempForIDQuery = `
		UPDATE users SET last_login_attempt = $1 WHERE id = $2
	`

	q := pgsqltx.QuerierFromCtx(ctx, i.db)
	r, err := q.ExecContext(ctx, updateUserLastLoginAttempForIDQuery, updatedDate, id)
	if err != nil {
		return errors.Wrap(err, "failed to update users tabler row")
	}

	if rowsAffected, err := r.RowsAffected(); err != nil {
		return errors.Wrap(err, "failed to get rows affected number")
	} else if rowsAffected == 0 {
		return errors.Wrap(user.ErrUserNotFound, "user not found. no records updated")
	}

	return nil
}
