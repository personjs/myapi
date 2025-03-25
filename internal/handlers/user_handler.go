// internal/handlers/user_handler.go
package handlers

import (
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User retrieved"))
}
