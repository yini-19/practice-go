package main

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"strings"
)

const maxInputLength = 500

var tmp, err = template.New("name").Parse("templates/index.html")

var allowed = map[string]bool{"standard": true, "shadow": true, "thinkertoy": true}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "wrong method", http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		http.Error(w, "page not found", http.StatusNotFound)
		return
	}

	var buf bytes.Buffer
	if err := tmp.ExecuteTemplate(&buf, "index.html", nil); err != nil {
		log.Printf("failed to execute template: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(buf.Bytes()); err != nil {
		log.Printf("failed to write response: %v", err)
	}
}

func AsciiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "wrong method", http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/ascii-art" {
		http.Error(w, "page not found", http.StatusNotFound)
		return
	}

	data := map[string]string{}
	contentType := r.Header.Get("Content-Type")
	if !strings.Contains(contentType, "application/x-www-form-urlencoded") {
		http.Error(w, "content type must be application/x-www-form-urlencoded", http.StatusBadRequest)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, "failed to parse form", http.StatusBadRequest)
		return
	}
	text, err := validateInput(r.PostForm, "text")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	banner, err := validateBanner(r.PostForm, "banner")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := generateAscii(banner, text)
	if err != nil {
		log.Printf("%v", err)
		http.Error(w, "failed to generate ASCII art", http.StatusInternalServerError)
		return
	}

	data["Message"] = "Welcome to Ascii"
	data["Result"] = result

	var buf bytes.Buffer
	if err := tmp.ExecuteTemplate(&buf, "index.html", nil); err != nil {
		log.Printf("failed to execute template: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(buf.Bytes()); err != nil {
		log.Printf("failed to write response: %v", err)
	}
}
