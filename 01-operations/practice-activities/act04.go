// Yi-Hsuan Hsu
// C277 Golang
// 6/28/2022

package main

import (
	"fmt"
)

// prompt for prinicipal, rate if interest, datys for loan
// calculate and simple interest for life of loan
func main() {
	var principal float64
	var interest_rate float64
	var days int

	fmt.Print("enter principal: ")
	fmt.Scanf("%f", &principal)
	fmt.Print("enter interest rate: ")
	fmt.Scanf("%f", &interest_rate)
	fmt.Print("enter days passed: ")
	fmt.Scanf("%d", &days)

	var interest float64 = principal * interest_rate * float64(days) / 365

	fmt.Printf("\ntotal interest is $%.2f\n", interest)

}
