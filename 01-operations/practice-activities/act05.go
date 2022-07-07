// Yi-Hsuan Hsu
// C277 Golang
// 6/28/2022

package main

import "fmt"

func main() {
	a, b := 5, 10
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)

	fmt.Printf("a < b ==> %v\n", a < b)
	fmt.Printf("a + 10 > b ==> %v\n", a+10 > b)
	fmt.Printf("b %% a == 0 ==> %v\n", b%a == 0)
	fmt.Printf("a == b ==> %v\n", a == b)
}
