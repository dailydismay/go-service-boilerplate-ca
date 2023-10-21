package tx

import (
	"context"
	"errors"
)

var (
	ErrUnexpetedTxFailure = errors.New("transaction unexpected failure")
)

type TransactionManager interface {
	Do(context.Context, func(context.Context) error) error
}
