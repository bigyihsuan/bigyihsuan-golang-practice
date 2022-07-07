// Yi-Hsuan Hsu
// C277 Golang
// 6/28/2022

package main

import (
	"fmt"
	"strings"
)

func main() {
	var season string

	fmt.Println("enter a season: ")
	fmt.Scan(&season)
	season = strings.ToLower(season)

	if season == "fall" || season == "autumn" {
		fmt.Println("I bet the leaves are pretty there!")
	} else if season == "winter" {
		fmt.Println("I hope you're ready for snow!")
	} else if season == "spring" {
		fmt.Println("I can smell the flowers!")
	} else if season == "summer" {
		fmt.Println("Make sure your AC is working!")
	} else {
		fmt.Println("I don't recognize that season.")
	}
}
