package main

import "fmt"

func main() {
	// ask for positive integer
	// display multiplication table from 1 to n^2
	var n uint
	fmt.Print("enter positive integer: ")
	fmt.Scanf("%d", &n)
	var i, j uint // why no types allowed in decl part of for?
	for i = 1; i <= n; i++ {
		for j = 1; j <= n; j++ {
			fmt.Printf("%3d ", i*j)
		}
		fmt.Println()
	}
}
