package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.back/internal/dto"
)

func (h *Handler) register(c *gin.Context) {
	var request dto.Register

	if err := c.BindJSON(&request); err != nil {
		logrus.Fatalln(err.Error())

		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	}

}

func (h *Handler) login(c *gin.Context) {

}
