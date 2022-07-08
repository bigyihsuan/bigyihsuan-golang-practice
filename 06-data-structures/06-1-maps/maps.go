package main

import (
	"fmt"
	"strings"
)

func main() {
	letters := make(map[string]int)

	value, isInMap := letters["k"]
	fmt.Println(value, isInMap)

	sentence := strings.ToLower("Lorem ipsum dolor sit amet, consectetur adipiscing elit")
	for _, c := range sentence {
		// fmt.Printf("before %c: %v\n", c, letters[string(c)])
		// fmt.Println(len(letters))
		if _, inMap := letters[string(c)]; !inMap {
			fmt.Printf("%c is not in map\n", c)
			letters[string(c)] = 1
		} else {
			fmt.Printf("%c in map\n", c)
			letters[string(c)] = 1
		}
		// fmt.Printf("after %c: %v\n", c, letters[string(c)])
		// fmt.Println(len(letters))
	}

}
// fmt.Println(letters)
// fmt.Println(letters["`z`"])
// for k, v := range letters {
// 	fmt.Printf("`%v` => %v\n", k, v)
// }
// fmt.Println()
// fmt.Println()

// var keys []string
// for k := range letters {
// 	keys = append(keys, k)
// }
// sort.Strings(keys)
// for _, k := range keys {
// 	fmt.Printf("`%v` => %v\n", k, letters[k])
// }
// fmt.Println()
