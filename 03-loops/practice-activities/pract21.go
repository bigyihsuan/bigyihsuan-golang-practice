package main

import "fmt"

func main() {
	// print armstrong numbers from 1 to 500: sum of each digit's cube is itself
	for i := 1; i <= 500; i++ {
		var sum = 0
		for j := i; j > 0; {
			digit := j % 10
			sum += digit * digit * digit
			j /= 10
		}
		if sum == i {
			fmt.Printf("%d is an armstrong number\n", i)
		}
	}
}
