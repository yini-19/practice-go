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
		renderError(w, "template parsing failed", 500)
		return
	}
	artist, err := fetchArtist("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		renderError(w, "failed to fetch artists", 500)
		return
	}

	if err := tmpl.Execute(w, artist); err != nil {
		renderError(w, "parsing failed", 500)
		return
	}
}

func artistDetailHandler(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("templates/artist.html")
	if err != nil {
		renderError(w, "template parsing failed", 500)
		return
	}

	path := r.URL.Path
	idstr := strings.TrimPrefix(path, "/artists/")

	id, err := strconv.Atoi(idstr)
	if err != nil {
		renderError(w, "parsing id failed", 400)
		return
	}

	artist, err := fetchArtist("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		renderError(w, "failed to fetch artists", 500)
		return
	}

	relation, err := fetchRelations("https://groupietrackers.herokuapp.com/api/relations")
	if err != nil {
		renderError(w, "failed to fetch relations", 500)
		return
	}
	if id == Relation.{

	}

	found := false
	for _, person := range artist {
		if id == person.ID {
			found = true
			if err := tmp.Execute(w, person); err != nil {
				renderError(w, "parsing failed", 500)
				return
			}
			break
		}

	}
	if !found {
		renderError(w, "no such artist", 404)
		return
	}
}
