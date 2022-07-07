// Yi-Hsuan Hsu
// C277 Golang
// 6/28/2022

package main

import "fmt"

func main() {
	var i1, i2, i3, i4, i5 int
	fmt.Print("enter 5 ints: ")
	fmt.Scanf("%d\n%d\n%d\n%d\n%d\n", &i1, &i2, &i3, &i4, &i5)

	var product = i1 * i2 * i3 * i4 * i5
	var sum = i1 + i2 + i3 + i4 + i5
	var average = float64(sum) / 5

	fmt.Printf("\nentered: %d %d %d %d %d\n", i1, i2, i3, i4, i5)
	fmt.Printf("product: %v\n", product)
	fmt.Printf("sum: %v\n", sum)
	fmt.Printf("average: %f\n", average)
	fmt.Println(4 + 7*5/3 + 2%10)
}
