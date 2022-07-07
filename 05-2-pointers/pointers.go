package main

import "fmt"

func main() {
	x := 34
	b := &x

	fmt.Printf("x: %v\n", x)
	fmt.Printf("*b: %v\n", *b)

	byVal(x)
	fmt.Printf("x: %v\n", x)
	fmt.Printf("*b: %v\n", *b)

	fmt.Printf("x: %v\n", x)
	fmt.Printf("*b: %v\n", *b)
	byAddr(b)
	fmt.Printf("x: %v\n", x)
	fmt.Printf("*b: %v\n", *b)

	// arr := [5]int{1, 2, 3, 4, 5}
	// fmt.Printf("before arr: %v\n", arr)
	// increment(arr[:])
	// fmt.Printf("after arr:  %v\n", arr)
}

func increment(arr []int) {
	for i := range arr {
		arr[i] += 1
	}
}

func byAddr(x *int) {
	*x = *x + 11
	fmt.Printf("byAddr *x: %v\n", *x)
}

func byVal(x int) {
	x = x + 11
	fmt.Printf("byVal x: %v\n", x)
}

// var arrptrs [5]*int

// for i := 0; i < len(arr); i++ {
// 	arrptrs[i] = &arr[i]
// }
// increment(arrptrs)
