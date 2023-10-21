package domainerr

import (
	"errors"
	"fmt"
)

// "github.com/pkg/errors"

var (
	rootErr = errors.New("domain defined error")
)

func New(msg string) error {
	return fmt.Errorf("%w: %v", rootErr, msg)
}

func Is(err error) bool {
	return errors.Is(err, rootErr)
}

func Join(errs ...error) error {
	return errors.Join(errs...)
}
