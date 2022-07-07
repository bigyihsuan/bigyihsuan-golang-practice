package main

import "fmt"

func main() {
	findDataType(10)
	findDataType(123.456)
	findDataType("hello")
	findDataType([4]int{1, 2, 3, 4})
	findDataType([]int{5, 6, 7, 8})
	findDataType(nil)
	findDataType('a')
	var i interface{}
	findDataType(i)
}

func findDataType(val interface{}) {
	switch t := val.(type) {
	case int:
		fmt.Printf("int %T => %T\n", t, t)
	case float32, float64:
		fmt.Printf("float %T => %T\n", t, t)
	case string:
		fmt.Printf("string %T => %T\n", t, t)
	case [4]int:
		fmt.Printf("array %T => %T\n", t, t)
	case []int:
		fmt.Printf("slice %T => %T\n", t, t)
	case nil:
		fmt.Printf("nil %T => %T\n", t, t)
	default:
		fmt.Printf("something else %T => %T\n", t, t)
	}
}
