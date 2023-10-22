package session

import "context"

type Repository interface {
	Save(context.Context, *Session) error
}
