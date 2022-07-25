package main

import (
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("./flatland01.txt")
	// data, err := ioutil.ReadFile("./flatland01.txt")
	file.Seek(-100, 2)
	text := make([]byte, 20)
	n, err := file.Read(text)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
	fmt.Println(string(text))
}
