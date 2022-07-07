// Yi-Hsuan Hsu
// C277 Golang
// 6/28/2022

package main

import "fmt"

// ask for how much money
// if > 20, rich
// otherwise broke

func main() {
	var wallet float64
	fmt.Print("what's in the wallet? ")
	fmt.Scanf("%f", &wallet)

	if wallet >= 20 {
		fmt.Println("you're rich")
	} else {
		fmt.Println("you're broke")
	}
}
