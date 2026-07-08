package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "parsing html failed", 500)
		return
	}
	if r.URL.Path != "/" {
		http.Error(w, "wrong path", 400)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "only GET method required", http.StatusMethodNotAllowed)
		return
	}
	tmpl.Execute(w, nil)
	// fmt.Fprintln(w, "This is our home page")
	// w.Write([]byte("my home page"))
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "parsing html failed", 500)
		return
	}

	if r.URL.Path != "/about" {
		http.Error(w, "wrong path", 404)
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "only POST allowed", 405)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	bannermap, err := LoadBanner("banners/" + banner + ".txt")
	if err != nil {
		http.Error(w, "banner loading failed", 404)
		return
	}

	asciiArt, err := GenerateArt(text, bannermap)
	if err != nil {
		http.Error(w, "parsing ascii failed", 500)
		return
	}
	
	
	tmpl.Execute(w, asciiArt)
}
