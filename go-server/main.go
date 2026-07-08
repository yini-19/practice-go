package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type Book struct {
	Title  string
	Author string
}

var books = []Book{
	{Title: "Go Programming", Author: "Alex Warri"},
	{Title: "Go Head First", Author: "Jerry West"},
	{Title: "Effective Coding", Author: "Casey Kings"},
}

type User struct {
	ID    int
	Name  string
	Email string
}

type User1 struct {
	Name string
	Age  int
}

func bookHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("books").Parse(` 
		<h1>Books</h1>
		<ul>
			{{range .}}
				<li>{{.Title}} - {{.Author}}</li>
			{{end}}
		</ul>
	`))

	tmpl.Execute(w, books)
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Query().Get("name")
	if user == "" {
		user = "Guest"
	}
	fmt.Fprintf(w, "Hello, %s!", user)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user := User{ID: 001, Name: "Adamawa Yola", Email: "adamawayola@gmail.com"}

	json.NewEncoder(w).Encode(user)

}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	var user User1
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Fprintf(w, "User %s registered successfully. Age: %d", user.Name, user.Age)
}
func main() {
	http.HandleFunc("POST /register", registerHandler)
	http.HandleFunc("/user", userHandler)
	http.HandleFunc("/greet", greetHandler)
	http.HandleFunc("/books", bookHandler)
	http.ListenAndServe(":8080", nil)
}
