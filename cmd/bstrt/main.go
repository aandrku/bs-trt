package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	//start a server
	err := http.ListenAndServeTLS(":8080", "tls/server.crt", "tls/server.key", mux)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
