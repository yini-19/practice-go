package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/type.html"))
	t.Execute(w, nil)
}

// type Person struct {
// 	Name string
// 	Age  int
// }

// var style = map[string]string{"block": "███", "line": "==="}

// func asciiLenght() {

// 	for num := 1; num <= 10; num++ {
// 		switch {
// 		case num <= 3:
// 			if num%2 == 0 {
// 				fmt.Println(num, "(even),(small)")
// 			} else {
// 				fmt.Println(num, "(odd),(small)")
// 			}
// 		case num > 3 && num < 7:
// 			if num%2 == 0 {
// 				fmt.Println(num, "(even),(medium)")
// 			} else {
// 				fmt.Println(num, "(odd),(medium)")
// 			}
// 		case num >= 7:
// 			if num%2 == 0 {
// 				fmt.Println(num, "(even),(large)")
// 			} else {
// 				fmt.Println(num, "(odd),(large)")
// 			}
// 		}

// 	}

// }
func main() {
	http.HandleFunc("/", Handler)
	fmt.Println("server starting on port 8080")
	http.ListenAndServe(":8080", nil)
	// person := Person{"Yiyakazah", 30}
	// var name = "Yiyakazah"
	// age := 30
	// fmt.Printf("Name: %s; Lenght: %d, Age: %d\n", name, len(name), age)
	// asciiLenght()
	// fmt.Println(person)
	// fmt.Println(style)
}
