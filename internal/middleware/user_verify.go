package middleware

import (
	"github.com/gin-gonic/gin"
	"go.back/pkg/customerror"
)

const autharizationHeader = "Authorization"

func UserVerify(c *gin.Context) {
	header := c.GetHeader(autharizationHeader)

	if header == "" {
		customerror.HandleHTTPError(c, customerror.Unauthorized)
	}
}
