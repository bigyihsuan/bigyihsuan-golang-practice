package main

import (
	"fmt"
)

type abc struct {
	a int
	b string
}

func main() {
	// Yi-Hsuan Hsu
	// name := "Robert Pike"
	// // find space
	// space := -1
	// for i, c := range name {
	// 	if unicode.IsSpace(c) {
	// 		space = i
	// 	}
	// }
	// // reverse each part
	// reversed := ""
	// for i := space - 1; i >= 0; i-- {
	// 	reversed += string(name[i])
	// }
	// reversed += " "
	// for i := len(name) - 1; i > space; i-- {
	// 	reversed += string(name[i])
	// }
	// fmt.Printf("%v\n%v\n", name, reversed)

	var k abc
	k.a = 10
	k.b = "hello"
	fmt.Printf("k: %#v\n", k)

	var copy = k
	copy.a = -1
	fmt.Printf("copy: %#v\n", copy)

	var reffed = &k
	reffed.a = 32
	fmt.Printf("reffed: %#v\n", reffed)
	fmt.Printf("k: %#v\n", k)
}

// s := ""
// for i := 'A'; i < 'E'; i++ {
// 	s += string(i)
// 	fmt.Println(s)
// }
// fmt.Println(s)

// for i := len(name) - 1; i >= 0; i-- {
// 	fmt.Printf("%3v    %3v    %v\n", i, name[i], string(name[i]))
// }
