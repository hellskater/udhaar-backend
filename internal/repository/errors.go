package repository

import (
	"errors"
)

var (
	// ErrNilID The ID of the general -purpose error argument is nil
	ErrNilID = errors.New("nil id")
	// ErrNotFound I can't find a general purpose error
	ErrNotFound = errors.New("not found")
	// ErrAlreadyExists General -purpose error already exists
	ErrAlreadyExists = errors.New("already exists")
	// ErrForbidden General -purpose error is prohibited
	ErrForbidden = errors.New("forbidden")
)

// ArgumentError Arrangement error
type ArgumentError struct {
	FieldName string
	Message   string
}

// Error Return Message
func (ae *ArgumentError) Error() string {
	return ae.Message
}

// ArgError Generate an argument error
func ArgError(field, message string) *ArgumentError {
	return &ArgumentError{FieldName: field, Message: message}
}

// IsArgError Is it an argument error?
func IsArgError(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(*ArgumentError)
	return ok
}
