package main

import "fmt"

func main() {
	var a int32 = 5
	var b int32 = 8
	var c int32 = 40
	fmt.Println(c + a - b)   // 40 + 5 - 8
	fmt.Println((c - a) * b) // (40 - 5) * 8
	fmt.Println(c / a / b)   // 40 / 5 / 8
	fmt.Println(c%a + b)     // 40%5 + 8
	fmt.Println((c % a) + b) // 40%5 + 8

}
