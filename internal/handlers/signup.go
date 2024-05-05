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
	http.ServeFile(w, r, "./templates/signup.html")

}

func PostSignup(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.FormValue("email")
	tag := r.FormValue("tag")
	password := r.FormValue("password")
	confirmPassword := r.FormValue("confirm-password")

	fmt.Println(email, tag, password, confirmPassword)

	//validate username
	if ok, _ := validate.Email(email); !ok {
		fmt.Println("Failed to validate email")
		http.Redirect(w, r, "/signup", http.StatusFound)
		return
	}

	//validate password
	if ok, _ := validate.Password(password); !ok || password != confirmPassword {
		fmt.Println("failed to validate password")
		http.Redirect(w, r, "/signup", http.StatusFound)
		return
	}

	//validate tag
	playerInfo, err := brawlapi.GetPlayerInfo(tag)
	if err != nil {
		fmt.Println("Failed to validate tag")
		http.Redirect(w, r, "/signup", http.StatusFound)
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
	_, err = db.Exec("INSERT INTO users (email, tag, hash) VALUES (?, ?, ?)", email, tag, hash)
	if err != nil {
		fmt.Println("failed to execute query")
		http.Redirect(w, r, "/signup", http.StatusFound)
		return
	}
	var id int
	err = db.QueryRow("SELECT id FROM users WHERE email = ?", email).Scan(&id)
	stmn, err := db.Prepare("INSERT INTO profiles (name, tag, club, icon, expLevel, threeVictories, duoVictories, soloVictories, user_id, trophies) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		fmt.Println("Failed to execute a prepare stmn")
	}
	defer stmn.Close()
	_, err = stmn.Exec(playerInfo.Name, playerInfo.Tag, playerInfo.Club.Tag, playerInfo.Icon.Id, playerInfo.ExpLevel, playerInfo.ThreeVSThreeVictories,
	playerInfo.DuoVictories, playerInfo.SoloVictories, id, playerInfo.Trophies)
	if err != nil {
		fmt.Println(err.Error())
	}
	
	//authenticate
	session, err := config.Store.Get(r, "login-session")
	if err != nil {
		log.Fatal(err)
	}
	session.Values["email"] = email 
	session.Values["authenticated"] = true

	//store in session
	session.Save(r, w)

	//send to the main page
	http.Redirect(w, r, "/", http.StatusFound)
	return

}
