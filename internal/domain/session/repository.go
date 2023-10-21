package session

import "context"

type Repository interface {
	FindByID(context.Context, *Session) error
	Save(context.Context, *Session) error
}
