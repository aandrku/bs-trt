package handlers

import (
	"bstrt/internal/config"
	"log"
	"net/http"
)

func GetLogin(w http.ResponseWriter, r *http.Request) {
	session, err := config.Store.Get(r, "login-session")
	if err != nil {
		log.Fatal(err)
	}

	if session.Values["authenticated"] == true {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	http.ServeFile(w, r, ".static/templates/login.html")
}
