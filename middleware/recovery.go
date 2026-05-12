package middleware

import (
	"net/http"

	"payment-service/response"

	"github.com/gin-gonic/gin"
)

// Recovery catches panics and returns a consistent 500 error response.
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				response.Error(c, http.StatusInternalServerError, "Internal server error", "an unexpected error occurred")
				c.Abort()
			}
		}()
		c.Next()
	}
}
