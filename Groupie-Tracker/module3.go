package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Groupie Tracker Home")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "About this project")
}

func artistHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/artists.html")
	if err != nil {
		http.Error(w, "template parsing failed", 500)
		return
	}
	artist, err := fetchArtist("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		http.Error(w, "failed to fetch artists", 500)
		return
	}

	if err := tmpl.Execute(w, artist); err != nil {
		http.Error(w, "parsing failed", 500)
		return
	}
}

func artistDetailHandler(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("templates/artist.html")
	if err != nil {
		http.Error(w, "template parsing failed", 500)
		return
	}
	path := r.URL.Path
	idstr := strings.TrimPrefix(path, "/artists/")

	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, "parsing id failed", 400)
		return
	}
	artist, err := fetchArtist("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		http.Error(w, "failed to fetch artists", 500)
		return
	}

	found := false
	for _, person := range artist {
		if id == person.ID {
			found = true
			if err := tmp.Execute(w, person); err != nil {
				http.Error(w, "parsing failed", 500)
				return
			}
			break
		}

	}
	if !found {
		http.Error(w, "ID not found!", 404)
		return
	}
}
