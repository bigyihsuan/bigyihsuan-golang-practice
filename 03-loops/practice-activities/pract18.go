package main

import (
	"fmt"
	"unicode"
)

func main() {
	// input string
	var s string = "hello 123  kjaskdljkalsjdklajdkl jlksada"

	// fmt.Print("input a string: ")
	// fmt.Scanln(&s)
	// fmt.Scanln()

	var numLetters, numSpaces, numDigits int

	for _, c := range s {
		// number of letters
		if unicode.IsLetter(c) {
			numLetters += 1
		}
		// number of spaces
		if unicode.IsSpace(c) {
			numSpaces += 1
		}
		// number of digits
		if unicode.IsDigit(c) {
			numDigits += 1
		}
	}
	fmt.Printf("numLetters: %v\n", numLetters)
	fmt.Printf("numSpaces: %v\n", numSpaces)
	fmt.Printf("numDigits: %v\n", numDigits)

}
