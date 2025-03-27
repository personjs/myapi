// routes/router.go
package routes

import (
	"personjs/myapi/internal/handlers"
	"personjs/myapi/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.New()

	// Manually add Logger and Recovery if needed
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Public route
	public := r.Group("/")
	public.GET("/", handlers.HomeHandler)
	public.GET("/health", handlers.HealthHandler)

	// Protected routes group
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware)
	protected.GET("/protected", handlers.ProtectedHandler)

	return r
}
