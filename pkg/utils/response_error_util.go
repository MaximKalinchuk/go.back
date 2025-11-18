package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message string `json:"message"`
}

func ResponseError(c *gin.Context, statusCode int, errorMessage string) {
	logrus.Errorln(errorMessage)
	c.AbortWithStatusJSON(statusCode, errorResponse{errorMessage})
}
