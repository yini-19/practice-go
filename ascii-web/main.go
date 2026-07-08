package main

import "net/http"

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-web", asciiHandler)
	http.HandleFunc("/ascii-web-switch", switchHandler)
	http.HandleFunc("/download", downloadHandler)
	http.ListenAndServe(":8080", nil)
}
