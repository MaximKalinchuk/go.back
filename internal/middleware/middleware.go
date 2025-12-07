package middleware

import "go.back/internal/service"

type Middleware struct {
	auth service.Authorization
}

func NewMiddleware(auth service.Authorization) *Middleware {
	return &Middleware{auth: auth}
}
