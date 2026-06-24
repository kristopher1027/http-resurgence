package main

import (
	"net/http"
	"fmt"
	"strconv"
)

func statushandler(w http.ResponseWriter,r *http.Request) {
	val := r.URL.Query().Get("code")
	if val == "" {
		http.Error(w,"bad request",http.StatusBadRequest)
		return
	}
	code, err := strconv.Atoi(val) 
	if err != nil {
		http.Error(w,"bad request",http.StatusBadRequest)
		return
	}
	if code < 100 || code > 599 {
		http.Error(w,"bad request",http.StatusBadRequest)
		return
	}
	statusText := http.StatusText(code)
	if statusText == "" {
		statusText = " " + statusText
	}
	 w.WriteHeader(code)
	fmt.Fprintf(w, "Responding with status %d%s", code, statusText)

}