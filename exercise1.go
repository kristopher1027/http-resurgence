package main

import (
	"net/http"
	"fmt"
)

func methodhandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	fmt.Fprintf(w,"You made a %s request.",r.Method)
}

