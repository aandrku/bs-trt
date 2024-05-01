package handlers

import (
	"bstrt/internal/config"
	"bstrt/internal/database"
	"bstrt/internal/util"
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

	http.ServeFile(w, r, "./static/templates/login.html")
}

func PostLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")

	db := database.ConnectionDB()
	defer db.Close()

	rows, err := db.Query("SELECT password FROM users WHERE username = ?", username)
	if err != nil {
		log.Fatal(err)
	}

	if !rows.Next() {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	var hash string

	err = rows.Scan(&hash)
	if err != nil {
		log.Fatal(err)
	}

	check := util.CheckPasswordHash(password, hash)

	if !check {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	session, err := config.Store.Get(r, "login-session")
	if err != nil {
		log.Fatal(err)
	}

	session.Values["username"] = username
	session.Values["authenticated"] = true
	session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusFound)
}
