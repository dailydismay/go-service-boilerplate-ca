package pgsql

import (
	"authsvc/internal/domain/session"

	"context"

	"github.com/jmoiron/sqlx"
)

type repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) session.Repository {
	return &repo{db}
}

/*
I simplified this part but I suggest placing an extra package QUERIES WITH:
- squirrel query builder
- database schema constants (such as tables)
*/
func (r *repo) FindByID(ctx context.Context, rt *session.Session) error {

	return nil
}

func (r *repo) Save(ctx context.Context, rt *session.Session) error {
	return nil
}
