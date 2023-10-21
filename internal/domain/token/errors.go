package token

import "authsvc/internal/core/domainerr"

var (
	ErrTokenExpired = domainerr.New("token expired")
	ErrInvalidSign  = domainerr.New("invalid sign")
)
