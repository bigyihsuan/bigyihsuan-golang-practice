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

	switch season {
	case "fall", "autumn":
		fmt.Println("I bet the leaves are pretty there!")
	case "winter":
		fmt.Println("I hope you're ready for snow!")
	case "spring":
		fmt.Println("I can smell the flowers!")
	case "summer":
		fmt.Println("Make sure your AC is working!")
	default:
		fmt.Println("I don't recognize that season.")
	}

}
