package customerror

import (
	"net/http"
)

type AppError struct {
	Code    int
	Message string
}

func (e *AppError) Error() string {
	return e.Message
}

var (
	UserNotFound       = &AppError{http.StatusNotFound, "Пользователь не найден"}
	InvalidCredentials = &AppError{http.StatusUnauthorized, "Неверный email или пароль"}
	Unauthorized       = &AppError{http.StatusUnauthorized, "Unauthorized"}
	ErrInternalServer  = &AppError{http.StatusInternalServerError, "Internal server error"}
)
