package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Health godoc
//
//	@Summary		Health check
//	@Description	Returns service health status
//	@Tags			system
//	@Produce		json
//	@Success		200	{object}	response.HealthResponse	"Service is healthy"
//	@Router			/health [get]
func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"service": "payment-service",
	})
}
