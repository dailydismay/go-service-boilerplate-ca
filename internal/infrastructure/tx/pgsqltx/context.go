package pgsqltx

import (
	"authsvc/internal/infrastructure/crosscutting/pgsql_client"
	"context"

	"github.com/jmoiron/sqlx"
)

type querierCtxKey struct{}

func QuerierFromCtx(ctx context.Context, db *sqlx.DB) pgsql_client.Querier {
	if tr := querierFromCtx(ctx); tr != nil {
		return tr
	}

	if db != nil {
		return pgsql_client.Querier(db)
	}

	return nil
}

func IsQuerierInContext(ctx context.Context) bool {
	return querierFromCtx(ctx) != nil
}

func (t *pgsqlTx) wrapContextAndGetCommitFn(ctx context.Context) (context.Context, commitFn, error) {
	if q := QuerierFromCtx(ctx, nil); q == nil {
		tx, err := t.db.Beginx()
		if err != nil {
			return nil, nil, err
		}

		return ctxWithQuerier(ctx, tx), sqlxCommit(tx), nil
	}

	return ctx, noopCommit(), nil
}

func querierFromCtx(ctx context.Context) pgsql_client.Querier {
	if tr, ok := ctx.Value(querierCtxKey{}).(pgsql_client.Querier); ok {
		return tr
	}

	return nil
}

func ctxWithQuerier(ctx context.Context, q pgsql_client.Querier) context.Context {
	return context.WithValue(ctx, querierCtxKey{}, q)
}
