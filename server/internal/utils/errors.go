package utils

import (
	"net/http"
)

type UserFacingError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *UserFacingError) Error() string {
	return e.Message
}

func NewUserFacingError(message string, code int) *UserFacingError {
	return &UserFacingError{
		Code:    code,
		Message: message,
	}
}

// Users
var (
	ErrInvalidCredentials = NewUserFacingError("invalid credentials", http.StatusForbidden)
	ErrUserDoNotExist     = NewUserFacingError("user doesn't exist", http.StatusNotFound)
	ErrPasswordDoNotMatch = NewUserFacingError("password do not match", http.StatusUnauthorized)
)
