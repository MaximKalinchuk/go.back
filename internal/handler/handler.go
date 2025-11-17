package handler

import (
	"github.com/gin-gonic/gin"
	"go.back/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("auth")
	{
		auth.POST("register", h.register)
		auth.POST("login", h.login)
	}

	return router
}
