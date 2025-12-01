package handler

import (
    "github.com/gin-gonic/gin"
    "go.back/pkg/customerror"
)

func HandleHTTPError(c *gin.Context, err error) {
    if appErr, ok := err.(*customerror.AppError); ok {
        c.AbortWithStatusJSON(appErr.Code, gin.H{
            "message": appErr.Message,
        })
        return
    }
    c.AbortWithStatusJSON(500, gin.H{"message": "Internal server error"})
}