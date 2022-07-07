package main

import "fmt"

// func A(x interface{}) {
// 	fmt.Println("calling A", x)
// }

// func B(x interface{}) {
// 	fmt.Println("calling B", x)
// }

// func C(x interface{}) {
// 	fmt.Println("calling C", x)
// }

func a() {
	fmt.Println("calling a")
}

func b(t int) {
	fmt.Println("calling b", t)
}

func c(s string, f float32) {
	fmt.Println("calling c", s, f)
}
func main() {
	// functions := []func(interface{}){A, B, C}
	// fmt.Printf("functions: %v\n", functions)
	// functions[0](10)
	// functions[1]("hello")
	// functions[2](5.34)

	f := []interface{}{a, b, c}

	f[0].(func())()
	f[1].(func(int))(10)
	f[2].(func(string, float32))("hello", 1.2345)
}
