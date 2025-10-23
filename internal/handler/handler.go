package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("api")
	{
		auth := api.Group("auth")
		{
			auth.POST("register", h.register)
			auth.POST("login", h.login)
		}
	}

	return router
}
