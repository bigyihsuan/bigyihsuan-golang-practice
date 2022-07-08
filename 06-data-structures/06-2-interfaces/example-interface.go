package main

import "fmt"

type animal interface {
	breathe()
	walk()
}

type human interface {
	animal
	speak()
}

type student struct {
	name string
}

func (s student) breathe() {
	fmt.Println("is breathing")
}
func (s student) walk() {
	fmt.Println("is walking")
}
func (s student) speak() {
	fmt.Println("is speaking")
}

type A interface {
	foo()
}
type B interface {
	foo()
}
type C struct {}

func (c C) foo() {
	// which foo()?
}


func main() {
	s := student {"bob"}
	s.breathe()
	s.walk()
	s.speak()
}