package main

import (
	"fmt"
	"math"
)

var p = fmt.Println

func main() {
	a := 3.0
	b := 1.3
	c := 7.5

	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("c: %v\n", c)

	p("a multiplied by c, with the result rounded up to the closest integer:\n\t", math.Ceil(a*c))
	p("a divided by b, with the result rounded down to the closest integer:\n\t", math.Floor(a/b))
	p("the square root of b, rounded to three decimal places (built-in):\n\t", math.Floor(math.Sqrt(b)*1000)/1000)
	p("the square root of b, rounded to three decimal places (user-defined):\n\t", roundToNthPlace(math.Sqrt(b), 3))
	p("c plus b, divided by a, with the result rounded up to the closest integer:\n\t", math.Ceil((c+b)/a))
	p("c squared (create two different statements for this value, using a different function in each):\n\t", math.Pow(c, 2))
	p("c squared (create two different statements for this value, using a different function in each):\n\t", c*c)
	p("identify the highest value of a, b, and c:\n\t", math.Max(a, math.Max(b, c)))
	p("identify the lowest value of a, b, and c:\n\t", math.Min(a, math.Min(b, c)))

}

func roundToNthPlace(x float64, n int) float64 {
	factor := math.Pow10(n)
	return math.Floor(x*factor) / factor
}
