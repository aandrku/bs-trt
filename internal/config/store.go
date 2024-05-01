package config

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

var Store *sessions.CookieStore

func init() {
	fmt.Println("started a cookie store")
	file, err := os.Open("./secret/sessions.key")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if scanner.Scan() {
		sk := scanner.Bytes()
		Store = sessions.NewCookieStore(sk)
		Store.Options = &sessions.Options{
			MaxAge:   3600 * 24, // Cookie expiration time in seconds (1 day)
			HttpOnly: true,      // HTTP only cookie
			Secure:   true,      // Cookie sent over HTTPS only
			SameSite: http.SameSiteDefaultMode,
		}
		fmt.Println("finished a cookie store")

		return
	}
	log.Fatalf("Failed to read a session-key")
}
