package middleware

import (
	"bstrt/internal/config"
	"log"
	"net/http"
)

func Auth(next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := config.Store.Get(r, "login-session")
		if err != nil {
			log.Fatal(err)
		}

		if session.IsNew || session.Values["authenticated"] != true {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		next(w, r)
	}
}
