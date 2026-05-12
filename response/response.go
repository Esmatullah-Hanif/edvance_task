package response

import "github.com/gin-gonic/gin"

// Success sends a 200 JSON envelope.
func Success(c *gin.Context, data any) {
	c.JSON(200, data)
}

// Error sends an error JSON envelope with the given HTTP status code.
func Error(c *gin.Context, statusCode int, message string, errDetail string) {
	c.JSON(statusCode, gin.H{
		"success": false,
		"message": message,
		"error":   errDetail,
	})
}
