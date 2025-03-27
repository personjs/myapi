package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Home handler function
func HomeHandler(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to the Home Route!")
}
