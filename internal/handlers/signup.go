package handlers

import (
	"bstrt/internal/brawlapi"
	"bstrt/internal/config"
	"bstrt/internal/database"
	"bstrt/internal/util"
	"bstrt/internal/validate"
	"fmt"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

func GetSignup(w http.ResponseWriter, r *http.Request) {
	//get a session from a request
	session, err := config.Store.Get(r, "login-session")
	if err != nil {
		log.Fatal(err)
	}

	//get a token from a session
	token := session.Values["jwt"]

	//if client logged in send him to the main page
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

	//serve a signup page
	http.ServeFile(w, r, "./templates/signup.html")
}

func PostSignup(w http.ResponseWriter, r *http.Request) {
	//parse a form
	r.ParseForm()
	//obtain values from a form
	email := r.FormValue("email")
	tag := r.FormValue("tag")
	password := r.FormValue("password")
	confirmPassword := r.FormValue("confirm-password")

	//validate username
	if ok, _ := validate.Email(email); !ok {
		SendVerificationError(w, r, "Invalid email.")
		return
	}

	//validate password
	fmt.Println(password, confirmPassword)
	if err := validate.Password(password); err != nil || password != confirmPassword {
		SendVerificationError(w, r, err.Error())
		return
	}

	//validate tag
	playerInfo, err := brawlapi.GetPlayerInfo(tag)
	if err != nil {
		SendVerificationError(w, r, "Invalid game tag.")
		return
	}

	//get db connection
	db := database.ConnectionDB()
	defer db.Close()

	//check if user already exists redirect to signup, if it does
	rows, err := db.Query("SELECT id FROM users WHERE email = ?", email)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	if rows.Next() {
		SendVerificationError(w, r, "User with such email already exists.")
		return
	}
	//hash Password
	hash, err := util.HashPassword(password)
	if err != nil {
		log.Fatal(err)
	}
	//store in a database
	_, err = db.Exec("INSERT INTO users (email, tag, hash) VALUES (?, ?, ?)", email, tag, hash)
	if err != nil {
		SendVerificationError(w, r, "Something went wrong.")
		return
	}
	var id int
	err = db.QueryRow("SELECT id FROM users WHERE email = ?", email).Scan(&id)
	stmn, err := db.Prepare("INSERT INTO profiles (name, tag, club, icon, expLevel, threeVictories, duoVictories, soloVictories, user_id, trophies) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		SendVerificationError(w, r, "Something went wrong.")
	}
	defer stmn.Close()
	_, err = stmn.Exec(playerInfo.Name, playerInfo.Tag, playerInfo.Club.Tag, playerInfo.Icon.Id, playerInfo.ExpLevel, playerInfo.ThreeVSThreeVictories,
		playerInfo.DuoVictories, playerInfo.SoloVictories, id, playerInfo.Trophies)
	if err != nil {
		SendVerificationError(w, r, "Something went wrong.")
	}

	//create a new jwt for a client
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

	//store jwt in a session
	session, err := config.Store.Get(r, "login-session")
	if err != nil {
		log.Fatal(err)
	}
	session.Values["jwt"] = encodedToken

	//store in session
	session.Save(r, w)

	//redirect to the main page
	fmt.Println("got to redirect")
	http.Redirect(w, r, "/", http.StatusFound)
	return
}
