package main

import (
	"ascii-art-web/ascii"
	"fmt"
	"html/template"
	"net/http"
)

// func about(w http.ResponseWriter, r *http.Request){
// 	w.Header()
// 	w.Write([]byte("About Page"))
// 	w.WriteHeader(http.StatusOK)
// }

func Handler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/index.html")
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return

	}

	data := map[string]string{
		"Message": "Welcome to ASCII-Art-Web!",
	}
	tmpl.Execute(w, data)

}
func AsciiHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/result.html")

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	asciiResult, err := ascii.GenerateAscii(text, banner)
	data := map[string]string{}

	if err != nil {
		data["Message"] = "Error: " + err.Error()
		data["Result"] = ""
	} else {
		//data["Message"] = "Here is your ascii art:"
		data["Result"] = asciiResult
	}
	tmpl.Execute(w, data)
	//fmt.Fprintln(w, r.URL.Path)//
}
func main() {
	http.HandleFunc("/", Handler)
	http.HandleFunc("/ascii-art-web", AsciiHandler)
	fmt.Println("server runing at http://localhost:8080")

	http.Handle("/statics/", http.StripPrefix("/statics/", http.FileServer(http.Dir("statics"))))
	http.ListenAndServe(":8080", nil)
}
