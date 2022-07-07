package main

import "fmt"

func main() {
	var cm float32
	const IN_TO_CM float32 = 2.54

	fmt.Print("Enter centimeters: ")
	fmt.Scanf("%f", &cm)

	var inches = cm / IN_TO_CM
	var feet = int(inches / 12)
	inches -= float32(12 * feet)

	fmt.Printf("%v cm is %v ft %v in\n", cm, feet, inches)
}
