// Yi-Hsuan Hsu
// C277 Golang
// 6/28/2022

package main

import (
	"fmt"
	"math"
)

// prompt user for flaot and return the in part of the float
func main() {
	fmt.Print("float number: ")
	var f float64
	// var fs string
	fmt.Scanf("%f", &f)
	fmt.Println("got", f)
	fmt.Println(math.Trunc(f))
}
