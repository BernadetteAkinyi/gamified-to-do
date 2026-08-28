package api

import (
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.Error(w, "URL Not Found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "index.html")

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
