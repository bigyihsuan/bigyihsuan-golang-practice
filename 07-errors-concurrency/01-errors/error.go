package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string = "10x"

	// the ParseInt function returns the parsed integer or
	// the error if the conversion failed
	a, error := strconv.ParseInt(s, 10, 8)
	fmt.Println(a)
	fmt.Println(error)
}
