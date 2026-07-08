package handlers

import (
	"ascii-art-web/internal/ascii"
	"net/http"
	"strings"
	"text/template"
)

var allowed = map[string]bool{"shadow": true, "standard": true, "thinkertoy": true}

func AsciiHandler(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("templates/result.html")
	if err != nil {
		http.Error(w, "Error parsing files", http.StatusInternalServerError)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "wrong method", http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/ascii-art" {
		http.Error(w, "page not found", http.StatusNotFound)
		return
	}
	data := map[string]string{}

	text := r.FormValue("input")
	text = strings.TrimSpace(text)
	if text == "" {
		http.Error(w, "empty input not allowed", http.StatusBadRequest)
		return
	}
	bannerFile := r.FormValue("banner")
	banner, err := ascii.BannerLoader("banners/" + bannerFile + ".txt")
	if err != nil {
		http.Error(w, "Error in loading banner file", http.StatusNotFound)
		return
	}
	asciiFunc := ascii.BuildArt(text, banner)
	data["Result"] = asciiFunc
	tmp.Execute(w, data)
}
