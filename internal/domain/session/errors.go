package session

import "errors"

var (
	ErrSessionExpired = errors.New("refresh token expired")
)
