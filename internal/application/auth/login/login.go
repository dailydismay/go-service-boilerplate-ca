package login

import (
	"authsvc/internal/core/hasher"
	"authsvc/internal/core/tx"
	"authsvc/internal/core/usecase"
	"authsvc/internal/domain/session"
	"authsvc/internal/domain/token"
	"authsvc/internal/domain/user"
)

type Payload struct {
	Username, Password string
}

type Result = token.GrantResult

type UseCase = usecase.UseCase[*Payload, *Result]

type implementation struct {
	txManager     tx.TransactionManager
	ph            hasher.Hasher
	userRepo      user.Repository
	tokenProvider token.Provider
	refreshRepo   session.Repository
}

func New(
	txManager tx.TransactionManager,
	ph hasher.Hasher,
	userRepo user.Repository,
	tokenProvider token.Provider,
	refreshRepo session.Repository,
) UseCase {
	return &implementation{txManager, ph, userRepo, tokenProvider, refreshRepo}
}
