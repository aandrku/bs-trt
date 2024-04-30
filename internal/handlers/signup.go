package handlers

import (
	"bstrt/internal/config"
	"bstrt/internal/database"
	"bstrt/internal/util"
	"bstrt/internal/validate"
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

func PostSignup(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")

	//validate username
	if ok, err := validate.Username(username); !ok {
		log.Print(err)
		http.Redirect(w, r, "/signup", http.StatusFound)
		return
	}

	//validate password
	if ok, err := validate.Password(password); !ok {
		log.Print(err)
		http.Redirect(w, r, "/signup", http.StatusFound)
		return
	}

	//get db connection
	db := database.ConnectionDB()
	defer db.Close()

	//check if user already exists redirect to signup, if it does
	rows, err := db.Query("SELECT id FROM users WHERE username = ?", username)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	if rows.Next() {
		log.Print("user already exists")
		http.Redirect(w, r, "/signup", http.StatusFound)
		return
	}
	//hash Password
	hash, err := util.HashPassword(password)
	if err != nil {
		log.Fatal(err)
	}
	//store in a database
	_, err = db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, hash)

	//authenticate
	session, err := config.Store.Get(r, "login-session")
	if err != nil {
		log.Fatal(err)
	}
	session.Values["username"] = username
	session.Values["authenticated"] = true

	//store in session
	session.Save(r, w)

	//send to the main page
	http.Redirect(w, r, "/", http.StatusFound)
	return

}
