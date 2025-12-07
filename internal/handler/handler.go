package handler

import (
	"github.com/gin-gonic/gin"
	"go.back/internal/middleware"
	"go.back/internal/service"
)

type Handler struct {
	services   *service.Service
	middleware *middleware.Middleware
}

func NewHandler(services *service.Service, middleware *middleware.Middleware) *Handler {
	return &Handler{
		services:   services,
		middleware: middleware,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(h.middleware.CORSMiddleware())

	auth := router.Group("auth")
	{
		auth.POST("register", h.register)
		auth.POST("login", h.login)
	}

	api := router.Group("api", h.middleware.UserVerify)

	users := api.Group("users")
	{
		users.GET(":id", h.getUser)
	}

	return router
}
