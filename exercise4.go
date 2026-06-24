package main
import (
	"net/http"
	"fmt"
)

func formhandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w,"method not allowed", http.StatusMethodNotAllowed)
		return
	}
	token := r.Header.Get("Content-Type")
	if token != "application/x-www-form-urlencoded" {
		http.Error(w, "Unsupported Media Type",415)
		return
	}

	if err := r.ParseForm();err != nil{
		http.Error(w,"bad request",http.StatusBadRequest)
		return
	}
	username := r.FormValue("username")
	language := r.FormValue("language")

	if username == "" {
		http.Error(w,"username is required",http.StatusBadRequest)
		return
	}
	if language == "" {
		http.Error(w,"language is required",http.StatusBadRequest)
		return
	}
fmt.Fprint(w, "Hello Ada, you are coding in Go!")
} 

