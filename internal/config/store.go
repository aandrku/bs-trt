package config

import (
	"bufio"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

var Store *sessions.CookieStore

func init() {
	file, err := os.Open("./secret/session.key")
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

		return
	}
	log.Fatalf("Failed to read a session-key")
}
