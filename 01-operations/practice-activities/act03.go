// Yi-Hsuan Hsu
// C277 Golang
// 6/28/2022

package main

import (
	"fmt"
	"math"
)

// calculate and display:
// current value of deposit for given initial deposit
// interest rate
// how many times interest is calculated per ear
// and the number of years since the initial deposit

// V = P(1 + r/n)^nt
// V -- value
// P -- initial deposit
// r -- interest rate as a fraction (eg 0.05)
// n -- the number of times per year interest is calculated
// t -- the number of years since the initial deposit

func main() {
	var initial_deposit float64
	var interest_rate float64
	var calcs_per_year int
	var years_passed int
	var current_value float64

	fmt.Print("Enter intial deposit ($): ")
	fmt.Scanf("%f", &initial_deposit)
	fmt.Print("Enter interest rate (%): ")
	fmt.Scanf("%f", &interest_rate)
	fmt.Print("Enter number of calculations per year: ")
	fmt.Scanf("%d", &calcs_per_year)
	fmt.Print("Enter number of years passed: ")
	fmt.Scanf("%d", &years_passed)

	current_value = initial_deposit * math.Pow((1+interest_rate/float64(calcs_per_year)), float64(calcs_per_year)*float64(years_passed))

	fmt.Printf("\ninitial value: $%.2f\n", initial_deposit)
	fmt.Printf("interest rate: %v%%\n", interest_rate)
	fmt.Printf("interest calculations per year: %v times\n", calcs_per_year)
	fmt.Printf("time passed: %v years\n", years_passed)
	fmt.Printf("current value: $%.2f\n", current_value)

}
