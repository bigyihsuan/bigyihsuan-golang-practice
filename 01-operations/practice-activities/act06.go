// Yi-Hsuan Hsu
// C277 Golang
// 6/28/2022

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Print("enter a number: ")
	var num int
	fmt.Scanf("%d", &num)

	fmt.Printf("\nselected %v\n", num)
	fmt.Printf("boolean of %v is %v\n", num, num != 0)
	fmt.Printf("binary of %v is %#b\n", num, num)
	fmt.Printf("sqrt of %v is %.3f\n", num, math.Sqrt(float64(num)))
}
