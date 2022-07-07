package main

import "fmt"

func main() {
	var message string = "HELLO WORLD"

	fmt.Println(message, len(message))

	for i, c := range message {
		fmt.Println(i, string(c))
	}

	// fmt.Print("number ")
	// var n int
	// fmt.Scanf("%d", &n)
	// sum := 0
	// for i := 0; i <= n; i++ {
	// 	sum += i
	// 	fmt.Println(sum)
	// }
}
