package main

import (
	"fmt"
	"strings"
)

func main() {
	var mapping map[string]int

	sentence := strings.ToLower("Lorem ipsum dolor sit amet, consectetur adipiscing elit")

	for i, c := range sentence {
		mapping[string(c)] = i
	}

	fmt.Println(mapping)
}
