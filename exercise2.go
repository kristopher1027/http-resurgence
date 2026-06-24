package main

import (
	"net/http"
	"fmt"
	"io"
)

func echohandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w,"Reject any non-POST request",http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	body,err := io.ReadAll(r.Body) 
		if err != nil {
			
		}
defer r.Body.Close()
		if len(body) == 0 {
			http.Error(w,"bad request",http.StatusBadRequest)
		}
		fmt.Fprint(w, "Hello Go")
	}
