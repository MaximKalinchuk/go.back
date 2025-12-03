package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getUser(c *gin.Context) {
	id := c.Param("id")

	user, err := h.services.User.GetUser(id)

	if err != nil {
		HandleHTTPError(c, err)
		return
	}

	c.JSON(http.StatusOK, user)

}
