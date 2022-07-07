package main

import "fmt"

func main() {
	arr := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(rowColSum(arr))
}

func rowColSum(arr [][]int) ([]int, []int) {
	rowSums := []int{}
	colSums := []int{}

	for i := 0; i < len(arr); i++ {
		rsum := 0
		csum := 0
		for j := 0; j < len(arr[i]); j++ {
			rsum += arr[i][j]
			csum += arr[j][i]
		}
		rowSums = append(rowSums, rsum)
		colSums = append(colSums, csum)
	}

	return rowSums, colSums
}

// fmt.Printf("arr: %v\n", arr)
// fmt.Printf("len(arr): %v\n", len(arr))
// fmt.Printf("arr[2]: %v\n", arr[2])
// fmt.Printf("len(arr[2]): %v\n", len(arr[2]))

// arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// other := make([]int, len(arr))
// fmt.Printf("other: %T %v\n", other, other)
// for i, e := range arr {
// 	other[len(arr)-i-1] = e
// }
// fmt.Printf("arr: %T %[1]v\n", arr)
// fmt.Println(len(arr))
// fmt.Printf("other: %T %[1]v\n", other)
