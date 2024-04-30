package handlers

import (
	"bstrt/internal/config"
	"log"
	"net/http"
)

func GetSignup(w http.ResponseWriter, r *http.Request) {
	//get session data
	session, err := config.Store.Get(r, "login-session")
	if err != nil {
		log.Fatal(err)
	}

	//if user logged in redirect him to the main page
	if session.Values["authenticated"] == true {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	//if user not logged in, send him a sign up page
	http.ServeFile(w, r, "./static/templates")

}
