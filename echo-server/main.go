package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/echo", EchoHandler)
	fmt.Println("server running at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}