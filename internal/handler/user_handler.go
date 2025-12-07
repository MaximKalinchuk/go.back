package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	userdto "go.back/internal/dto/user"
	"go.back/pkg/customerror"
)

func (h *Handler) getUser(c *gin.Context) {
	userId := c.GetString("userId")

	if userId == "" {
		customerror.HandleHTTPError(c, customerror.Unauthorized)
		return
	}

	user, err := h.services.User.GetUser(userId)

	if err != nil {
		customerror.HandleHTTPError(c, err)
		return
	}

	c.JSON(http.StatusOK, userdto.FromEntity(&user))

}
