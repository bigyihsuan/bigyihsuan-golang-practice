// Yi-Hsuan Hsu
// C277 Golang
// 6/28/2022

package main

import (
	"fmt"
	"strings"
)

func main() {
	var hasCats, hasDogs bool

	var response string
	fmt.Print("have cats? ")
	fmt.Scan("%s", &response)
	hasCats = strings.ToLower(response) == "yes"

	fmt.Print("have dogs? ")
	fmt.Scan("%s", &response)
	hasDogs = strings.ToLower(response) == "yes"

	if hasCats && hasDogs {
		fmt.Println("You must really love pets!")
	} else {
		fmt.Println("Maybe you need more pets.")
	}
}
