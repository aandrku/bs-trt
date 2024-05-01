package handlers

import "net/http"

func GetRoot(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/templates/index.html")
}
