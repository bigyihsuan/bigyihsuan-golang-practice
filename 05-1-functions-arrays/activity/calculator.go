//! Yi-Hsuan Hsu
//! 6/30/22
//! C277 Golang
//! Activity: Calculator

package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	for {
		var as, bs, ops string

		fmt.Print("Enter left: ")
		fmt.Scan(&as)
		if isExit(as) {
			break
		}
		fmt.Print("Enter right: ")
		fmt.Scan(&bs)
		if isExit(bs) {
			break
		}

		a, err_a := strconv.ParseFloat(as, 64)
		b, err_b := strconv.ParseFloat(bs, 64)
		if err_a != nil || err_b != nil {
			fmt.Printf("Please enter valid numbers (got `%v` and `%v`)\n", as, bs)
			continue
		}

		fmt.Print("Enter an operation\n(add, sub, mul, div, mod, sqrt, fact): ")
		fmt.Scan(&ops)
		if isExit(ops) {
			break
		}

		switch ops {
		case "add":
			fmt.Printf("%v + %v = %v", a, b, add(a, b))
		case "sub":
			fmt.Printf("%v - %v = %v", a, b, sub(a, b))
		case "mul":
			fmt.Printf("%v * %v = %v", a, b, mul(a, b))
		case "div":
			if b == 0 {
				fmt.Println("ERROR: division by 0, try again")
				continue
			}
			fmt.Printf("%v / %v = %v", a, b, div(a, b))
		case "mod":
			if b == 0 {
				fmt.Println("ERROR: division by 0, try again")
				continue
			}
			fmt.Printf("%v %% %v = %v", a, b, mod(a, b))
		case "sqrt":
			if a < 0 {
				fmt.Println("ERROR: negative sqrt, try again")
				continue
			}
			fmt.Printf("sqrt(%v) = %v", a, sqrt(a))
		case "fact":
			if a < 0 {
				fmt.Println("ERROR: negative factorial, try again")
				continue
			}
			fmt.Printf("%v! = %v", a, fact(a))
		}
		fmt.Println()
	}
}

func isExit(s string) bool {
	return strings.ToLower(s) == "quit" || strings.ToLower(s) == "exit"
}

func add(a, b float64) float64 {
	return a + b
}
func sub(a, b float64) float64 {
	return a - b
}
func mul(a, b float64) float64 {
	return a * b
}
func div(a, b float64) float64 {
	return a / b
}
func mod(a, b float64) float64 {
	return float64(int(a) % int(b))
}
func sqrt(a float64) float64 {
	return math.Sqrt(a)
}
func fact(a float64) float64 {
	if a <= 2 {
		return float64(int(a))
	}
	return float64(int(a) * int(fact(a-1)))
}
