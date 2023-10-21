package token

import (
	"authsvc/internal/domain/user"
	"context"
)

type GrantResult struct {
	Access, Refresh Token
}

type Provider interface {
	Grant(context.Context, user.ID) (*GrantResult, error)
}
