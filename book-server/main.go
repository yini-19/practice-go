package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Book struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	PublishedYear int    `json:"published_year"`
	Publisher     string `json:"publisher"`
	Genre         string `json:"genre"`
	ISBN          string `json:"isbn"`
	Language      string `json:"language"`
}

var bookstore = map[string]Book{
	"Book1": {ID: "uuid-1",
		Title:         "Go Programming Made Simple",
		Author:        "Yiyakazah Nicodemus",
		PublishedYear: 2026,
		Publisher:     "CodeBase Publishing",
		Genre:         "Technology",
		ISBN:          "1234567898760",
		Language:      "English"},

	"Book2": {ID: "uuid-2",
		Title:         "Go Programming for Dummies",
		Author:        "Lantana Yusuf",
		PublishedYear: 2026,
		Publisher:     "CodeBase Publishing",
		Genre:         "Technology",
		ISBN:          "123456789567890",
		Language:      "English"},
}

func bookHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	books := make([]Book, 0, len(bookstore))
	for _, b := range bookstore {
		books = append(books, b)
	}
	if err := json.NewEncoder(w).Encode(books); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getBookHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	parts := strings.Split(r.URL.Path, "/")

	if parts[1] != "books" || len(parts) != 3 {
		http.NotFound(w, r)
		return
	}
	id := parts[2]

	book, found := bookstore[id]
	if !found {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Book not Found"})
		return
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(book)
	}

}
func main() {
	http.HandleFunc("/books", bookHandler)
	http.HandleFunc("/books/", getBookHandler)
	fmt.Println("server running at http://localhost:8080/books/")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("error")
	}
}
/*
package main

import (
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/google/uuid"
)

type Book struct {
    ID            string `json:"id"`
    Title         string `json:"title"`
    Author        string `json:"author"`
    PublishedYear int    `json:"publishedyear"`
    Publisher     string `json:"publisher"`
    Genre         string `json:"genre"`
    ISBN          string `json:"isbn"`
    Language      string `json:"language"`
}

var bookstore = make(map[string]Book)

// POST /books handler
func createBookHandler(w http.ResponseWriter, r *http.Request) {
    var book Book

    // Decode JSON body
    if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Generate UUID for the book ID
    book.ID = uuid.New().String()

    // Store in map
    bookstore[book.ID] = book

    // Return created book with 201 Created
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(book)
}

func main() {
    http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodGet {
            // GET /books
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(bookstore)
        } else if r.Method == http.MethodPost {
            // POST /books
            createBookHandler(w, r)
        } else {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    })

    fmt.Println("Server running at http://localhost:8080/books")
    http.ListenAndServe(":8080", nil)
}
*/