package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomNums(size, max int) []int {
	numbers := make([]int, size)
	for i := 0; i < size; i++ {
		numbers[i] = rand.Intn(max + 1)
	}
	return numbers
}

func sort(arr []int, ascending bool) {
	// insertion sort
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if !ascending && arr[i] > arr[j] {
				swap(arr, i, j)
			} else if ascending && arr[i] < arr[j] {
				swap(arr, i, j)
			}
		}
	}
}

func sortMerge(left []int, right []int, ascending bool) []int {
	var merged []int
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if ascending {
			if left[i] <= right[j] {
				merged = append(merged, left[i])
				i++
			} else {
				merged = append(merged, right[j])
				j++
			}
		} else {
			if left[i] >= right[j] {
				merged = append(merged, left[i])
				i++
			} else {
				merged = append(merged, right[j])
				j++
			}

		}
		// fmt.Printf("i: %v\n", i)
		// fmt.Printf("j: %v\n", j)
		// fmt.Printf("merged: %v\n", merged)
	}
	// fmt.Printf("i: %v\n", i)
	// fmt.Printf("j: %v\n", j)
	// fmt.Printf("merged: %v\n", merged)
	if i < len(left) {
		merged = append(merged, left[i:]...)
	} else {
		merged = append(merged, right[j:]...)
	}
	return merged
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func main() {
	rand.Seed(time.Now().Unix())
	left := randomNums(10, 100)
	right := randomNums(10, 100)

	fmt.Printf("left: %v\n", left)
	fmt.Printf("right: %v\n", right)

	sort(left, true)
	sort(right, true)

	fmt.Printf("sorted left: %v\n", left)
	fmt.Printf("sorted right: %v\n", right)

	fmt.Printf("sortMerge(left, right, ascending): %v\n", sortMerge(left, right, true))
}
