package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	user, pass, hasAuth := c.Request.BasicAuth()
	if !hasAuth || user != "admin" || pass != "admin" {
		c.Header("WWW-Authenticate", `Basic realm="Restricted"`)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	c.Next()
}
