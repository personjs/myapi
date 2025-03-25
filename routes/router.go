// routes/router.go
package routes

import (
	"personjs/myapi/internal/handlers"
	"personjs/myapi/pkg/middleware"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	// Apply AuthMiddleware globally to all routes
	r.Use(middleware.AuthMiddleware)

	// Define routes
	r.HandleFunc("/health", handlers.HealthCheck).Methods("GET")
	r.HandleFunc("/user", handlers.GetUser).Methods("GET")

	return r
}
