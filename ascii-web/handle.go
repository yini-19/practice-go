package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type PageData struct {
	Result    string
	Banner    string
	Text      string
	Linecount int
	Charcount int
}

var tmp = template.Must(template.ParseFiles("templates/index.html"))

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "incorrect path", 400)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "only GET method is supported", http.StatusMethodNotAllowed)
		return
	}

	if err := tmp.Execute(w, PageData{}); err != nil {
		http.Error(w, "HTML parsing failed", 500)
		return
	}
}

func asciiHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-web" {
		http.Error(w, "incorrect path", 400)
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "only POST method is supported", http.StatusMethodNotAllowed)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if text == "" || banner == "" {
		http.Error(w, "field cannot be empty", 400)
		return
	}

	for _, char := range text {
		if char == '\r' || char == '\n' {
			continue
		}
		if char < rune(32) || char > rune(126) {
			http.Error(w, "only printable ascii characters allowed", 400)
			return
		}
	}

	bannermap, err := LoadBanner("banners/" + banner + ".txt")
	if err != nil {
		http.Error(w, "error loading banner file", 500)
		return
	}

	asciiArt, err := GenerateArt(text, bannermap)
	if err != nil {
		http.Error(w, "failed to generate art", 500)
		return
	}

	linecount := strings.Count(asciiArt, "\n")
	charcount := len(text)

	result := PageData{
		Result: asciiArt,
		Banner: banner,
		Text:   text,
		Linecount: linecount,
		Charcount: charcount,
	}
	if err := tmp.Execute(w, result); err != nil {
		http.Error(w, "HTML parsing failed", 500)
		return
	}
}

func switchHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-web-switch" {
		http.Error(w, "incorrect path", 400)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "only GET method is supported", http.StatusMethodNotAllowed)
		return
	}

	text := r.URL.Query().Get("text")
	banner := r.URL.Query().Get("banner")

	if text == "" || banner == "" {
		http.Error(w, "field cannot be empty", 400)
		return
	}

	for _, char := range text {
		if char == '\r' || char == '\n' {
			continue
		}
		if char < rune(32) || char > rune(126) {
			http.Error(w, "only printable ascii characters allowed", 400)
			return
		}
	}

	bannermap, err := LoadBanner("banners/" + banner + ".txt")
	if err != nil {
		http.Error(w, "error loading banner file", 500)
		return
	}

	asciiArt, err := GenerateArt(text, bannermap)
	if err != nil {
		http.Error(w, "failed to generate art", 500)
		return
	}

	linecount := strings.Count(asciiArt, "\n")
	charcount := len(text)

	result := PageData{
		Result: asciiArt,
		Banner: banner,
		Text:   text,
		Linecount: linecount,
		Charcount: charcount,
	}
	if err := tmp.Execute(w, result); err != nil {
		http.Error(w, "HTML parsing failed", 500)
		return
	}
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/download" {
		http.Error(w, "incorrect path", 400)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	text := r.URL.Query().Get("text")
	banner := r.URL.Query().Get("banner")

	if text == "" || banner == "" {
		http.Error(w, "field cannot be empty", 400)
		return
	}

	for _, char := range text {
		if char == '\r' || char == '\n' {
			continue
		}
		if char < rune(32) || char > rune(126) {
			http.Error(w, "only printable ascii characters allowed", 400)
			return
		}
	}

	bannermap, err := LoadBanner("banners/" + banner + ".txt")
	if err != nil {
		http.Error(w, "error loading banner file", 500)
		return
	}

	asciiArt, err := GenerateArt(text, bannermap)
	if err != nil {
		http.Error(w, "failed to generate art", 500)
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Disposition", `attachment; filename="ascii-art.txt"`)
	fmt.Fprint(w, asciiArt)
}
