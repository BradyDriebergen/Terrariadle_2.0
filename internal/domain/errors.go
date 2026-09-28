package domain

import "errors"

// errors/errors.go (or domain/errors.go)
type ErrCode string

const (
	ErrUsrNotFound  ErrCode = "USER_NOT_FOUND"
	ErrNotFound     ErrCode = "NOT_FOUND"
	ErrConflict     ErrCode = "CONFLICT"
	ErrInvalidInput ErrCode = "INVALID_INPUT"
	ErrInternal     ErrCode = "INTERNAL"
)

type AppError struct {
	Code    ErrCode
	Message string
	Err     error
}

func (e *AppError) Error() string { return e.Message }
func (e *AppError) Unwrap() error { return e.Err }

// Constructors
func UserNotFound(msg string, cause error) *AppError {
	return &AppError{Code: ErrUsrNotFound, Message: msg, Err: cause}
}

func NotFound(msg string, cause error) *AppError {
	return &AppError{Code: ErrNotFound, Message: msg, Err: cause}
}

func Conflict(msg string, cause error) *AppError {
	return &AppError{Code: ErrConflict, Message: msg, Err: cause}
}

func InvalidInput(msg string, cause error) *AppError {
	return &AppError{Code: ErrInvalidInput, Message: msg, Err: cause}
}

func Internal(msg string, cause error) *AppError {
	return &AppError{Code: ErrInternal, Message: msg, Err: cause}
}

// Error used for if a document doesn't exist in MongoDB
var MongoErrNotFound = errors.New("not found")
