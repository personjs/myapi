package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthHandler returns a simple health check response
func HealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "Service is healthy",
	})
}
