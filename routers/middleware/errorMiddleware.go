package middleware

import (
	"hito/routers/resp"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		lastErr := c.Errors.Last()
		if lastErr == nil {
			c.Next()
			return
		}

		if apiError, ok := errors.Cause(lastErr.Err).(resp.Error); ok {
			c.JSON(apiError.HttpStatusCode, gin.H{"code": apiError.Code, "message": apiError.Message})
			return
		}

		c.Next()
	}
}
