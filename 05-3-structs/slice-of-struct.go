package main

import "fmt"

type class struct {
	ClassName string
	students  []student
}

type student struct {
	Name   string
	Number int
	City   string
}

func main() {
	bob := student{Name: "bob", Number: 1, City: "NYC"}
	dan := student{Name: "dan", Number: 6, City: "Los Angeles"}

	students := []student{bob, dan, {Name: "mary", Number: 3, City: "Austin"}}

	// fmt.Printf("students: %v\n", students)

	class := class{ClassName: "science", students: students}
	fmt.Printf("class: %v\n", class)

	i := []interface{}{1, 2, "def", 3}
	j := []interface{}{"abc", "hij"}

	k := append(i, j...)
	fmt.Printf("k: %v\n", k)
}
