// cmd/main.go
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"personjs/myapi/config"
	"personjs/myapi/routes"
)

func main() {
	// Define flags
	hostPtr := flag.String("h", "localhost", "Host to run the server on")
	portPtr := flag.String("p", "8080", "Port to run the server on")

	// Parse flags
	flag.Parse()

	// Load configuration
	config.LoadConfig()

	// Set host and port from arguments
	host := *hostPtr
	port := *portPtr
	address := fmt.Sprintf("%s:%s", host, port)

	// Setup the router
	r := routes.SetupRouter()

	// Start server
	fmt.Println("Server running on", address)
	log.Fatal(http.ListenAndServe(address, r))
}
