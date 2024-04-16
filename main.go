package main

import (
	"log"
	"net/http"
	"prueba_tecnica_golang/user-services/handlers"
)

func main() {
	// Create a new ServeMux (router)
	mux := http.NewServeMux()

	// Register handlers
	mux.HandleFunc("/register", handlers.RegisterUser)
	mux.HandleFunc("/login", handlers.LoginUser)

	// Set up the HTTP server configuration
	server := &http.Server{
		Addr:    ":8080", // Define the port to listen on
		Handler: mux,     // Set the router as the handler
	}

	log.Println("Starting server on port 8080...")
	// Start the HTTP server
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
