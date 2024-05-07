package middleware

import (
	"bstrt/internal/config"
	"bstrt/internal/util"
	"log"
	"net/http"
)

func Auth(next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := config.Store.Get(r, "login-session")
		if err != nil {
			log.Fatal(err)
		}

		if session.IsNew {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		token := session.Values["jwt"].(string)
		_, claims, err := util.ParseToken(token, []byte("secret"))

		if claims["authenticated"] != true {
			http.Redirect(w, r, "/login", http.StatusFound)
		}

		next(w, r)
	}
}
