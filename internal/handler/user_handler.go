package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	userdto "go.back/internal/dto/user"
	"go.back/pkg/customerror"
)

func (h *Handler) getUser(c *gin.Context) {
	id := c.Param("id")

	user, err := h.services.User.GetUser(id)

	if err != nil {
		customerror.HandleHTTPError(c, err)
		return
	}

	c.JSON(http.StatusOK, userdto.FromEntity(&user))

}
