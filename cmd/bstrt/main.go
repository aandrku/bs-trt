package main

import (
	"bstrt/internal/handlers"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /signup", handlers.GetSignup)

	//start a server
	err := http.ListenAndServeTLS(":8080", "tls/server.crt", "tls/server.key", mux)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
