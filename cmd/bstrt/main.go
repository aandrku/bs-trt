package main

import (
	"bstrt/internal/handlers"
	"bstrt/internal/middleware"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", middleware.Auth(handlers.GetRoot))
	mux.HandleFunc("GET /signup", handlers.GetSignup)
	mux.HandleFunc("POST /signup", handlers.PostSignup)
	mux.HandleFunc("GET /login", handlers.GetLogin)
	mux.HandleFunc("POST /login", handlers.PostLogin)

	fs := http.FileServer(http.Dir("static"))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fs))

	//start a server
	err := http.ListenAndServeTLS(":8080", "tls/cert.pem", "tls/key.pem", mux)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
