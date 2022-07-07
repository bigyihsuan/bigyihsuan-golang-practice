package main

import "fmt"

func main() {
	x, y, z := 1, 10, 5

	if x > y && y > z {
		fmt.Println("x greater than z")
	} else if y > z {
		fmt.Println("y greater than z")

	} else {
		fmt.Println("z is biggest")
	}

	y = 19
	switch y {
	case 4, 5:
		fmt.Println("4 or 5")
	default:
		fmt.Println("default")
	}
}
