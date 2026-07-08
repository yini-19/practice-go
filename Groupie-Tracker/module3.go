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

func artistDetailHandler(w http.ResponseWriter, r *http.Request){
	path := r.URL.Path
	path = strings.TrimPrefix(path, "/artists/")
	
	idstr, err := strconv.Atoi(path)
	if err != nil {
		http.Error(w, "parsing id failed", 400)
		return
	}
	artist, err := fetchArtist("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		http.Error(w, "failed to fetch artists", 500)
		return
	}

	for _, person := range artist{
		if person.ID == idstr {

		}
	
	}
}