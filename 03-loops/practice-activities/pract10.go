package main

import "fmt"

func main() {
	// take a number, produce a number triangle
	var n int = 0
	fmt.Print("enter n: ")
	fmt.Scanf("%d", &n)
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d", i)
		}
		fmt.Println()
	}

}
