package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	dto "go.back/internal/dto/auth"
	"go.back/pkg/customerror"
)

func (h *Handler) register(c *gin.Context) {
	var request dto.Register

	if err := c.BindJSON(&request); err != nil {
		customerror.HandleHTTPError(c, err)
		return
	}

	userId, err := h.services.Authorization.CreateUser(request)

	if err != nil {
		customerror.HandleHTTPError(c, err)
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": userId,
	})
}

func (h *Handler) login(c *gin.Context) {
	var requset dto.Login

	if err := c.BindJSON(&requset); err != nil {
		customerror.HandleHTTPError(c, err)
		return
	}

	token, err := h.services.Authorization.GenerateToken(requset)

	if err != nil {
		customerror.HandleHTTPError(c, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
