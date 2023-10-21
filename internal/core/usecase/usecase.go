package usecase

import "context"

type UseCase[P any, R any] interface {
	Execute(context.Context, P) (R, error)
}

type Interactor[P any] interface {
	Execute(context.Context, P) error
}
