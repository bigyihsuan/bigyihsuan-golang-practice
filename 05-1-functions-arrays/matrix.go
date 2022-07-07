package main

import "fmt"

func main() {
	var matrix [][]int = makeMatrix(3, 3)

	// for i := range matrix {
	// 	matrix[i] = getThreeInputs(3)
	// }

	for i, row := range matrix {
		for j := range row {
			// fmt.Printf("int element [%v, %v]: ", i, j)
			// var in int
			// fmt.Scanf("%v", &in)
			// matrix[i][j] = in
			matrix[i][j] = i*len(matrix) + j + 1
		}
	}

	// fmt.Println(matrix)
	// printBottomTriangle(matrix)
	// fmt.Println(switchFirstAndLastCols(matrix))
	// fmt.Println(switchFirstAndLastRows(matrix))

	thing := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Printf("thing: %v\n", thing)
	fmt.Print("pop(thing, 3): ")
	fmt.Println(pop(thing, 3))

	thing = []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Printf("reverse(thing): %v\n", reverse(thing))

	thing = []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Printf("thing: %v\n", thing)
	fmt.Println(mapArr(thing, func(i int) int { return i * 2 }))
	fmt.Println(filter(thing, func(i int) bool { return i%2 != 0 }))
	fmt.Printf("thing: %v\n", thing)
	fmt.Println(reduce(thing, func(r, v int) int { return r + v }))
}

func pop(arr []int, x int) (int, []int) {
	n := arr[x]
	return n, append(arr[:x], arr[x+1:]...)[:len(arr)-1]
}

func reverse(arr []int) []int {
	var newArr []int
	for i := len(arr) - 1; i >= 0; i-- {
		newArr = append(newArr, arr[i])
	}
	return arr
}

func mapArr(arr []int, f func(int) int) []int {
	var newArr []int
	for _, v := range arr {
		newArr = append(newArr, f(v))
	}
	return arr
}

func reduce(arr []int, f func(int, int) int) int {
	var result int
	for _, v := range arr {
		result = f(result, v)
	}
	return result
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

func makeMatrix(x, y int) [][]int {
	mat := make([][]int, x)
	for i := range mat {
		mat[i] = make([]int, y)
	}
	return mat
}

func getThreeInputs(size int) []int {
	var row []int = make([]int, size)
	for i := range row {
		fmt.Printf("int element [%v]: ", i)
		var in int
		fmt.Scanf("%v", &in)
		row[i] = in
	}
	return row
}

func printBottomTriangle(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("%v ", matrix[i][j])
		}
		fmt.Println()
	}
}

func switchFirstAndLastCols(matrix [][]int) [][]int {
	newMat := makeMatrix(len(matrix), len(matrix[0]))
	for i, row := range matrix {
		for j, k := len(row)-1, 0; j >= 0; j-- {
			newMat[i][j] = row[k]
			k++
		}
	}
	return newMat
}

func switchFirstAndLastRows(matrix [][]int) [][]int {
	newMat := makeMatrix(len(matrix), len(matrix[0]))
	for i, k := len(matrix)-1, 0; i >= 0; i-- {
		newMat[i] = matrix[k]
		k++
	}
	return newMat
}
