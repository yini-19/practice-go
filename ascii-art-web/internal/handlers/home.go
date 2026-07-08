package handlers

import (
	"bytes"
	"log"
	"net/http"
	"text/template"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Error parsing files", http.StatusInternalServerError)
		return
	}

	if r.Method != http.MethodGet {
		renderError(w, http.StatusMethodNotAllowed,"wrong method")
		return
	}
	if r.URL.Path != "/" {
		renderError(w, http.StatusNotFound,"page not found")
		return
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, nil); err != nil {
		renderError(w,http.StatusInternalServerError, "error writing")
		return
	}
	if _, err := w.Write(buf.Bytes()); err != nil {
		log.Printf("failed to write file: %v", err)
	}
}
