package main

import "fmt"

func main() {
	// frequency of each digit in a given int
	fmt.Print("input n: ")
	var n int64
	fmt.Scanf("%d", &n)

	var digits = [10]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for n > 0 {
		digits[n%10]++
		n /= 10
	}

	for digit, count := range digits {
		if count <= 0 {
			continue
		}
		if count == 1 {
			fmt.Printf("%d occurs %d time\n", digit, count)
		} else {
			fmt.Printf("%d occurs %d times\n", digit, count)
		}
	}

}
