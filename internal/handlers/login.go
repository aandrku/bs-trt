package handlers

import (
	"bstrt/internal/config"
	"bstrt/internal/database"
	"bstrt/internal/util"
	"fmt"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

func GetLogin(w http.ResponseWriter, r *http.Request) {
	//get a session
	session, err := config.Store.Get(r, "login-session")
	if err != nil {
		log.Fatal(err)
	}

	//get a token value from a session
	token := session.Values["jwt"]

	// if user is authenticated send them to a main page
	if !session.IsNew && token != nil {
		token := token.(string)
		_, claims, err := util.ParseToken(token, []byte("secret"))
		if err != nil {
			fmt.Println(err.Error())
		}
		if claims["authenticated"] == true {
			http.Redirect(w, r, "/", http.StatusFound)
		}
	}

	//serve a login page
	http.ServeFile(w, r, "./templates/login.html")
}

func PostLogin(w http.ResponseWriter, r *http.Request) {
	//parse a form from a request
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err.Error())
	}

	//get an email and password values from a form
	email := r.FormValue("email")
	password := r.FormValue("password")

	//establish database connection
	db := database.ConnectionDB()
	defer db.Close()

	//get hash of user
	rows, err := db.Query("SELECT hash FROM users WHERE email = ?", email)
	if err != nil {
		log.Fatal(err)
	}

	//if database didnt't return any rows - users doesnt exist return an error to a user
	if !rows.Next() {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	//put hash into a varible
	var hash string
	err = rows.Scan(&hash)
	if err != nil {
		log.Fatal(err)
	}

	//check if password match
	check := util.CheckPasswordHash(password, hash)
	if !check {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	//create a new jwt and store it in clients session
	claims := jwt.MapClaims{
		"email":         email,
		"authenticated": true,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := "secret"

	encodedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		fmt.Println(err.Error())
	}

	session, err := config.Store.Get(r, "login-session")
	if err != nil {
		log.Fatal(err)
	}

	session.Values["jwt"] = encodedToken
	session.Save(r, w)

	//redirect client to a main page
	http.Redirect(w, r, "/", http.StatusFound)
}
