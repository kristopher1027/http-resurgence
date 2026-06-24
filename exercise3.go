package main

import (
	"net/http"
	"fmt"
)

func headerhandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	token := r.Header.Get("X-Custom-Token")
	if token == "" {
		http.Error(w, "X-Custom-Token header is missing", http.StatusBadRequest)
		return
	}
	contentType := r.Header.Get("Content-Type")
	if contentType == "" {
		contentType = "Content-Type not provided"
	}
	fmt.Fprintf(w, "Token received: %s\nContent-Type: %s", token, contentType)
}






// 	w.Header().Set("Content-Type", "text/plain")

// 	word := r.Header.Get("X-Custom-Token") 
// 	if word == "" {
// 		http.Error(w,"X-Custom-Token header is missing",http.StatusBadRequest)
// 		return
// 	}
// 	token := r.Header.Get("X-Custom-Type")
// 	if token == "" {
// 		http.Error(w,"Content-Type not provided",http.StatusBadRequest)
// 		return

// 	}
		

// 	fmt.Fprintf(w,"word received: %s\nContent-Type: %s",word,token)
// }