//! Yi-Hsuan Hsu
//! 6/30/22
//! C277 Golang
//! Assessment: Fizz Buzz

package main

import "fmt"

func main() {
	// Ask the user for a number.
	var n int
	fmt.Print("How many fizzing and buzzing units do you need in your life? ")
	fmt.Scan(&n)
	// Output a count starting with 0.
	fmt.Println(0)
	for i, fizzBuzzCount := 1, 0; fizzBuzzCount < n; i++ {
		// Display the count number if it is not divisible by 3 or 5.
		// Replace every multiple of 3 with the word "fizz."
		// Replace every multiple of 5 with the word "buzz."
		// Replace multiples of both 3 and 5 with "fizz buzz."
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("fizz buzz")
			fizzBuzzCount++
		case i%3 == 0:
			fmt.Println("fizz")
			fizzBuzzCount++
		case i%5 == 0:
			fmt.Println("buzz")
			fizzBuzzCount++
		default:
			fmt.Println(i)
		}
		// Continue counting until the number of integers replaced with "fizz," "buzz," or "fizz buzz" reaches the input number.
	}
	// The last output line should read "TRADITION!!"
	fmt.Println("TRADITION!!")
}
