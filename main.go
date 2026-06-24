package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/method-inspector", methodhandler)
	http.HandleFunc("/echo", echohandler)
	http.HandleFunc("/headers", headerhandler)
	http.HandleFunc("/form", formhandler)
	http.HandleFunc("/status", statushandler)
	fmt.Println("server is running http://localhost:8080")
	http.ListenAndServe(":8080",nil)
}