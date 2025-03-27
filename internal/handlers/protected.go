// internal/handlers/user_handler.go
package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ProtectedHandler serves a protected route
func ProtectedHandler(c *gin.Context) {
	c.String(http.StatusOK, "You have accessed a protected route!")
}
