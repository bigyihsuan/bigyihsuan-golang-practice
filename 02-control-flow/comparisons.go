package main

import "fmt"

func main() {
	// do not change the variable declarations
	var a int8 = 100 // architect-specific type
	var b int = 70   // implementation-specific type

	fmt.Println("a =", a)
	fmt.Println("b =", b)
	fmt.Println("a == b:", int(a) == b) //checks for equality
	fmt.Println("a != b:", int(a) != b) //checks for inequality
}
