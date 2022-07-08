package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Create a function that takes as input an int n and returns an array of length n with random integers between -100 and 100.
func randoms(n int) []int {
	rand.Seed(time.Now().Unix())
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		numbers[i] = rand.Intn(201) - 100 // (0, 200) - 100 = (-100, 100)
	}
	return numbers
}

// Compute the max of an array of int
// Compute the index of the max of an array of int
func max(arr []int) (int, int) {
	largest := -1
	index := -1
	for i := range arr {
		if arr[i] > largest {
			largest = arr[i]
			index = i
		}
	}
	return largest, index
}

// Compute the min of an array of int
// Compute the index of the min of an array of int
func min(arr []int) (int, int) {
	smallest := -1
	index := -1
	for i := range arr {
		if arr[i] < smallest {
			smallest = arr[i]
			index = i
		}
	}
	return smallest, index
}

// Sort an array of int in descending order and return the new sorted array in a separate array.
func sortedDesc(arr []int) []int {
	var sorted []int
	sorted = append(sorted, arr...)
	// insertion sort
	for i := 0; i < len(sorted); i++ {
		for j := 0; j < len(sorted); j++ {
			if sorted[i] > sorted[j] {
				swap(sorted, i, j)
			}
		}
	}
	return sorted
}

// Sort an array of int in ascending order and return the new sorted array in a separate array.
func sortedAsc(arr []int) []int {
	var sorted []int
	sorted = append(sorted, arr...)
	// insertion sort
	for i := 0; i < len(sorted); i++ {
		for j := 0; j < len(sorted); j++ {
			if sorted[i] <= sorted[j] {
				swap(sorted, i, j)
			}
		}
	}
	return sorted
}

// Compute the mean of an array
func mean(arr []int) int {
	return sum(arr) / len(arr)
}

// Compute the median of an array
func median(arr []int) int {
	sorted := sortedAsc(arr)
	return sorted[len(sorted)/2]
}

// Identify all positive numbers in the array and return the numbers in a slice
func pos(arr []int) []int {
	return filter(arr, isPositive)
}

// Identify all negative numbers in the array and return the numbers in a slice
func neg(arr []int) []int {
	return filter(arr, isNegative)
}

// Compute the longest sequence of sorted numbers (in descending or ascending order) in the array and return in a new array
// 		Example: input: [1 45 67 87 6 57 0]
// 		Output: [1 45 67 87]
func longestSorted(arr []int) []int {
	// iterate through the arr
	// keep track of what direction the current run is
	// if the current element doesn't match the current direction, reset the direction and its slice
	// at the end return the longest slice
	var desc, asc, longestdesc, longestasc []int
	desc = append(desc, arr[0])
	asc = append(asc, arr[0])
	direction := 0 // 0 = not sorted, 1 = ascending, -1 = descending

	// fmt.Printf("asc: %v\n", asc)
	// fmt.Printf("desc: %v\n", desc)
	// fmt.Printf("longestasc: %v\n", longestasc)
	// fmt.Printf("longestdesc: %v\n", longestdesc)
	// fmt.Println()
	for i := 1; i < len(arr); i++ {
		switch direction {
		case 0:
			if desc[len(desc)-1] > arr[i] {
				desc = append(desc, arr[i])
				asc = []int{arr[i]}
				direction = -1
			} else {
				asc = append(asc, arr[i])
				desc = []int{arr[i]}
				direction = 1
			}
		case -1: // looking at a descending run
			if desc[len(desc)-1] > arr[i] {
				desc = append(desc, arr[i])
			} else {
				// reset if opposite direction
				desc = []int{arr[i]}
				asc = append(asc, arr[i])
				direction = 1
			}
		case 1: // looking at an ascending run
			if asc[len(asc)-1] < arr[i] {
				asc = append(asc, arr[i])
			} else {
				// reset if opposite direction
				asc = []int{arr[i]}
				desc = append(desc, arr[i])
				direction = -1
			}
		}
		if len(longestdesc) < len(desc) {
			longestdesc = desc
		}
		if len(longestasc) < len(asc) {
			longestasc = asc
		}
		// fmt.Printf("current: %v\n", arr[i])
		// fmt.Printf("asc: %v\n", asc)
		// fmt.Printf("desc: %v\n", desc)
		// fmt.Printf("longestasc: %v\n", longestasc)
		// fmt.Printf("longestdesc: %v\n", longestdesc)
		// fmt.Println()
	}
	if len(longestdesc) > len(longestasc) {
		return longestdesc
	} else {
		return longestasc
	}
}

// Remove duplicates from an array of ints and return the unique elements in a slice
func unique(arr []int) []int {
	unique := make(map[int]int)
	for _, v := range arr {
		unique[v] = v
	}
	return keys(unique)
}

// helpers
func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func sum(arr []int) int {
	s := 0
	for _, v := range arr {
		s += v
	}
	return s
}

func filter(arr []int, f func(int) bool) []int {
	var newArr []int
	for _, v := range arr {
		if f(v) {
			newArr = append(newArr, v)
		}
	}
	return newArr
}
func isPositive(n int) bool {
	return n > 0
}
func isNegative(n int) bool {
	return n < 0
}
func isSortedDesc(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < arr[i+1] {
			return false
		}
	}
	return true
}
func isSortedAsc(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

func keys(m map[int]int) []int {
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func main() {
	numbers := randoms(100)
	fmt.Printf("numbers: %v\n", numbers)

	min, mini := min(numbers)
	fmt.Printf("min(numbers): %v @ %v\n", min, mini)
	max, maxi := max(numbers)
	fmt.Printf("max(numbers): %v @ %v\n", max, maxi)

	fmt.Printf("sortedDesc(numbers): %v\n", sortedDesc(numbers))
	fmt.Printf("sortedAsc(numbers): %v\n", sortedAsc(numbers))

	fmt.Printf("mean(numbers): %v\n", mean(numbers))
	fmt.Printf("median(numbers): %v\n", median(numbers))

	fmt.Printf("pos(numbers): %v\n", pos(numbers))
	fmt.Printf("neg(numbers): %v\n", neg(numbers))

	fmt.Printf("longestSorted(numbers): %v\n", longestSorted(numbers))
	fmt.Printf("unique(numbers): %v\n", unique(numbers))
}
