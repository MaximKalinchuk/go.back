package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.back/internal/dto"
	"go.back/pkg/utils"
)

func (h *Handler) register(c *gin.Context) {
	var request dto.Register

	if err := c.BindJSON(&request); err != nil {
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := h.services.Authorization.CreateUser(request)

	if err != nil {
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": userId,
	})
}

func (h *Handler) login(c *gin.Context) {

}
