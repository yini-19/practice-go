package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong")
}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	user := r.URL.Query().Get("name")
	if user == "" {
		user = "Guest"
	}
	fmt.Fprintf(w, "Hello, %v!", user)
}

func countHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Send a POST request with text to count words", http.StatusMethodNotAllowed)
		return
	}
	text, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "error reeadeing text", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "the length of inputed text is %q", len(text))
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	operation := r.URL.Query().Get("op")
	num1:= r.URL.Query().Get("a")
	num2 := r.URL.Query().Get("b")

	number1, err := strconv.Atoi(num1)
	if err != nil {
		http.Error(w, "parsing failed", http.StatusBadRequest)
		return
	}
	number2, err := strconv.Atoi(num2)
	if err != nil {
		http.Error(w, "parsing failed", http.StatusBadRequest)
		return
	}

	switch operation {
	case "add":
		fmt.Fprintf(w, "Result: %v", number1+number2)
		return
	case "subtract":
		fmt.Fprintf(w, "Result: %v", number1-number2)
		return
	case "multiply":
		fmt.Fprintf(w, "Result: %v", number1*number2)
		return
	default:
		http.Error(w, "operation is unknown", http.StatusBadRequest)
		return
	}

}
func dashboardHandler(w http.ResponseWriter, r *http.Request){
	APIKey := "secret123"
	apikey := r.Header.Get("X-API-Key")

	if apikey != APIKey{
		http.Error(w, "API key not recognised", http.StatusUnauthorized)
		return
	}
	fmt.Fprint(w, "Welcome, authorisation successful")
	
}

func agentHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, r.Header)
}

func legacyHandler(w http.ResponseWriter, r *http.Request){
	http.Redirect(w, r, "/v2", http.StatusMovedPermanently)
}

func v2Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w,"Welcome to version 2")
}

func main() {
	http.HandleFunc("/v2", v2Handler)
	http.HandleFunc("/legacy", legacyHandler)
	http.HandleFunc("/agent", agentHandler)
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/count", countHandler)
	http.HandleFunc("/calculate", calculateHandler)
	http.HandleFunc("/dashboard", dashboardHandler)
	http.ListenAndServe(":8080", nil)
}
