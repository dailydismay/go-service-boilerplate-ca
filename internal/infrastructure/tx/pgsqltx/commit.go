package pgsqltx

import (
	"authsvc/internal/core/tx"
	"errors"

	"github.com/jmoiron/sqlx"
)

type commitFn func(*error) error

func sqlxCommit(sqlxTx *sqlx.Tx) commitFn {
	return func(errP *error) error {
		if p := recover(); p != nil {
			_ = sqlxTx.Rollback()
			panic(p)
		}

		if *errP != nil {
			if rerr := sqlxTx.Rollback(); rerr != nil {
				return errors.Join(*errP, tx.ErrUnexpetedTxFailure, rerr)
			}
		}

		if err := sqlxTx.Commit(); err != nil {
			return errors.Join(tx.ErrUnexpetedTxFailure, err)
		}

		return nil
	}
}

func noopCommit() commitFn {
	return func(err *error) error {
		if *err != nil {
			return *err
		}

		return nil
	}
}
