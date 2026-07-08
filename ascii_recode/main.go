package main

import (
	"fmt"
)

func main() {

	fmt.Println(LoadBanner("standard.txt"))
	fmt.Println(LoadBanner("shadow.txt"))
	fmt.Println(LoadBanner("notfound.txt"))
}
