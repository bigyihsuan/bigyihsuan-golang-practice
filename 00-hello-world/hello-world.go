package main

import (
	"fmt"
)

func main() {
	var name string = "bigyihsuan"
	var address string = "123 main street"
	var birth_year int = 1999

	fmt.Print("enter address: ")
	// fmt.Scan(&address)

	// var birth_year_str string
	// fmt.Print("enter birth year: ")
	// fmt.Scan(&birth_year_str)
	// birth_year, _ = strconv.Atoi(birth_year_str)
	// fmt.Println()

	fmt.Scanln(&address)

	fmt.Printf("%s lives on %s and was born on %d\n", name, address, birth_year)
	fmt.Printf("name: %p\naddress: %p\nbirth_year: %p\n", &name, &address, &birth_year)
	fmt.Printf("name: %T\naddress: %T\nbirth_year: %T\n", name, address, birth_year)
}
