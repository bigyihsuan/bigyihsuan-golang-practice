package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// generate a random int between 0 and 10
	// user guesses
	// if high, too high
	// if low, too low
	// if same, that's right
	// repeat until the number is guessed

	rand.Seed(time.Now().UnixNano())
	var number = rand.Intn(11)
	var guess = -1
	fmt.Print("guess number\n> ")
	for guess != number {
		fmt.Scanf("%d", &guess)
		if guess > number {
			fmt.Println("Too high, try again.")
			fmt.Print("> ")
		} else if guess < number {
			fmt.Println("Too low, try again.")
			fmt.Print("> ")
		} else {
			fmt.Println("That's right. You guessed the number!")
		}
	}
}
