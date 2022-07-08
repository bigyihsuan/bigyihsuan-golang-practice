package main

/*
It prompts the user for a series of 5-10 integers.
The user must be prompted for a minimum of 5 numbers.
After the user has entered at least 5 numbers, the program should allow the user to stop entry or enter another number.
When the user chooses to stop or after 10 numbers have been entered, the program stops prompting for values and performs the following calculations:
the product of the integers
the sum of the integers
After performing the calculations, the program should display the following to the user:
the values the user entered
each of the calculations, using a phrase that identifies the value
*/

import (
	"fmt"
	"strings"
)

func add(r, v int) int {
	return r + v
}

func mult(r, v int) int {
	return r * v
}

func reduce(arr []int, start int, f func(int, int) int) int {
	var result int = start
	for _, v := range arr {
		result = f(result, v)
	}
	return result
}

func main() {
	var numbers []int

	for i := 0; i < 5; i++ {
		var num int
		fmt.Print("enter a number: ")
		fmt.Scan(&num)
		numbers = append(numbers, num)

	}
	for {
		var response string
		fmt.Print("stop entry? (y/n) ")
		fmt.Scan(&response)
		response = strings.ToLower(response)
		if response == "y" {
			fmt.Println("no more numbers")
			break
		}
		var num int
		fmt.Print("enter a number: ")
		fmt.Scan(&num)
		numbers = append(numbers, num)
		if len(numbers) >= 10 {
			fmt.Println("no more numbers")
			break
		}
	}

	sum := reduce(numbers, 0, add)
	prod := reduce(numbers, 1, mult)

	fmt.Printf("numbers entered: %v\n", numbers)
	fmt.Printf("sum of the numbers: %v\n", sum)
	fmt.Printf("product of the numbers: %v\n", prod)
}
