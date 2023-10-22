package common

import (
	"errors"
	"net/http"
)

type ErrorDetail struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type HttpError struct {
	root   error
	Status int
	// TODO: map usecase errors to client-server error dictionary
	Code    string        `json:"code"`
	Message string        `json:"message"`
	Details []ErrorDetail `json:"details"`
}

func (he *HttpError) Error() string {
	return he.root.Error()
}

type HttpErrorBuilder struct {
	e *HttpError
}

func ErrorBuilder(err error) *HttpErrorBuilder {
	e := &HttpError{
		root:    err,
		Status:  http.StatusBadRequest,
		Message: err.Error(),
		Details: make([]ErrorDetail, 0),
	}
	b := &HttpErrorBuilder{e}

	return b
}

func (b *HttpErrorBuilder) Build() error {
	return b.e
}

func (b *HttpErrorBuilder) Plain() *HttpError {
	return b.e
}

func (b *HttpErrorBuilder) Status(status int) *HttpErrorBuilder {
	b.e.Status = status
	return b
}

func (b *HttpErrorBuilder) Detail(k, v string) *HttpErrorBuilder {
	b.e.Details = append(b.e.Details, ErrorDetail{k, v})
	return b
}

func AsHttpError(err error) (*HttpError, bool) {
	var e *HttpError
	return e, errors.As(err, &e)
}
