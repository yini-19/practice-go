package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var tmpl = template.Must(template.ParseFiles("templates/index.html"))

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Wrong Method", http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		http.Error(w, "wrong path", http.StatusNotFound)
		return
	}
	tmpl.Execute(w, nil)
}

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/echo" {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}
	text := r.FormValue("text")
	if text == "" {
		http.Error(w, "wrong request", http.StatusBadRequest)
	}
	//tmpl.Execute(w, text)
	fmt.Fprintln(w, text)
}
