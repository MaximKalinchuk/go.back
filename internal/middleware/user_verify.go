package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"go.back/pkg/customerror"
)

const autharizationHeader = "Authorization"

func (m *Middleware) UserVerify(c *gin.Context) {
	header := c.GetHeader(autharizationHeader)

	if header == "" {
		customerror.HandleHTTPError(c, customerror.Unauthorized)
		return
	}

	tokenParts := strings.Split(header, " ")

	if len(tokenParts) != 2 {
		customerror.HandleHTTPError(c, customerror.Unauthorized)
		return
	}

	userId, err := m.auth.ParseToken(tokenParts[1])

	if err != nil {
		customerror.HandleHTTPError(c, customerror.Unauthorized)
		return
	}

	c.Set("userId", userId)
}
