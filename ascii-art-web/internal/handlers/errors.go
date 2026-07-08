package handlers

import (
	"html/template"
	"net/http"
)

type ErrorData struct {
	Code    int
	Message string
}

var templs = template.Must(template.ParseFiles("templates/error.html"))

func renderError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)

	data := ErrorData{
		Code:    code,
		Message: message,
	}
	err := templs.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}