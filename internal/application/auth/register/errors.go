package register

import "authsvc/internal/core/domainerr"

var (
	ErrAccountAlreadyExists = domainerr.New("account already exists")
)
