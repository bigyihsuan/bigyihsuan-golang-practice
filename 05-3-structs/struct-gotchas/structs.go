package main

import (
	"fmt"
	"time"
	"unsafe"

	"struct-gotchas/secretagent"
	"struct-gotchas/treesort"
)

// type StructName struct {
type Employee struct {
	ID            int
	Name, Address string // you can combine field declaration of the same type like this
	DoB           time.Time
	Position      string
	Salary        int
	ManagerID     int
}

var employees []*Employee

func main() {
	var bill Employee

	// Access struct fields with a dot.
	bill.ID = 1
	bill.Name = "Bill"
	bill.Address = "123 Main Street"
	bill.DoB = time.Now()
	bill.Position = "Boss"
	bill.Salary = 5000
	bill.ManagerID = 0

	// Can define with a literal with values in the order declared in the struct.
	peter := Employee{
		2,
		"Peter",
		"10 Downing Street",
		time.Now(),
		"Programmer",
		1000,
		1,
	}

	// Can also use named fields to only specify the fields we want.
	// The rest will be zero-valued.
	micheal := Employee{
		Name:     "Micheal",
		Position: "Programmer",
		Salary:   1000,
	}

	// The zero value of a struct is composed of the
	// zero values of all its fields.
	var milton Employee
	fmt.Printf("milton: %#v\n", milton)
	// When designing a struct, it's a good idea to
	// define a zero value that is easy to use.
	// For example, bytes.Buffer is a ready-to-use empty buffer struct.

	// Can modify struct fields with dot notation.
	peter.Salary -= 50 // didn't come in on Saturday or Sunday

	// Can take the address of fields
	position := &peter.Position
	*position = "Senior " + *position

	// And also the struct itself
	straightShooter := &peter
	straightShooter.Position += " (upper management written all over him)"

	// Notice the dot notation works for both struct pointers and struct values.
	// These lines are equivalent:
	// straightShooter.Position += " (upper management written all over him)"
	// (*straightShooter).Position += " (upper management written all over him)"

	fmt.Printf("bill: %#v\n", bill)
	fmt.Printf("peter: %#v\n", peter)
	fmt.Printf("micheal: %#v\n", micheal)

	employees = append(employees, &bill, &peter, &micheal)

	// If the return type of EmployeeByID were an Employee instead of a pointer to one
	// this would not compile, because the left side would not acutally "live anywhere".
	// We would need to assign it first to a variable first.
	// pete := EmployeeByID(0) // necessary if we return a value instead of a pointer
	// pete.Salary += 5000
	EmployeeByID(0).Salary += 5000
	fmt.Printf("EmployeeByID(0): %v\n", EmployeeByID(0))

	// Structs are comparable if all their fields are comparable (therefore can be used as keys in maps).
	samir := Employee{
		Name: "Samir",
	}
	samir_clone := Employee{
		Name: "Samir",
	}
	fmt.Printf("samir == samir_clone: %t\n", samir == samir_clone)

	type address struct {
		hostname string
		port     int
	}
	hits := make(map[address]int)
	hits[address{"golang.org", 443}]++
	fmt.Printf("hits: %v\n", hits)

	// We can make anonymous structs by declaring them without a name and using a struct literal
	// immediately after the type.
	// anakin := struct {
	// 	Name string
	// 	Rank string
	// }{
	// 	Name: "Anakin",
	// 	Rank: "Jedi (not a master)",
	// }
	// The fields of structs live contiguously in memory in the order
	// they were defined.
	// If you are defining anonymous structs, you must
	// make sure their fields are defined in the same order if you want to compare them.
	// obiwan := struct {
	// 	Rank string
	// 	Name string
	// }{
	// 	Rank: "Jedi Master",
	// 	Name: "Obiwan",
	// }
	// fmt.Printf("anakin == obiwan: %t\n", anakin == obiwan) // this will not compile because the order of the fields is not the same

	// Only capitalized struct fields are exported from a package.
	// See the secretagent package.
	agent := secretagent.InitAgent(7)
	// agent.doubleO = false             // doubleO is undefined to this package
	fmt.Printf("agent: %#v\n", agent) // but we can see it when we print...
	// So what if we try to initialize a struct with the order-based literal syntax?
	// After all, we're not using the name...
	// agent2 := secretagent.Agent{6, true}
	// fmt.Printf("agent2: %#v\n", agent2) // no. `implicit assignment to unexported field doubleO in secretagent.Agent literal`

	// structs can have recursively contain themselves,
	// but only as pointers.
	// See the treesort package for an example.
	s := []int{7, 2, 1, 3, 5, 6, 4}
	treesort.Sort(s)
	fmt.Printf("s: %#v\n", s)

	// The struct type with no fields is called the empty struct.
	// Why would you use it?
	// Well it has zero size, so you can use it as placeholder.
	// There are actually several niche uses for it, beyond the scope of this lesson.
	// An article to meditate on, should you like to know more: https://dave.cheney.net/2014/03/25/the-empty-struct
	var x [10000]struct{}
	var y [10000]bool

	fmt.Printf("x size: %#v\n", unsafe.Sizeof(x))
	fmt.Printf("y size: %#v\n", unsafe.Sizeof(y))

	// We can create structs within structs.
	// type Point struct {
	// 	X, Y int
	// }
	// type Circle struct {
	// 	Center Point
	// 	Radius int
	// }
	// type Wheel struct {
	// 	Circle Circle
	// 	Spokes int
	// }
	// However, this syntax is cumbersome.
	// See how we have to dig into the struct to get to the fields.
	// var w Wheel
	// w.Circle.Center.X = 8
	// w.Circle.Center.Y = 8
	// w.Circle.Radius = 5
	// w.Spokes = 20
	// A better way is use this syntax to "embed" a struct within a struct.
	type Point struct {
		X, Y int
	}
	type Circle struct {
		Point
		Radius int
	}
	type Wheel struct {
		Circle
		Spokes int
	}
	var w Wheel
	w.X = 8 // equivalent to w.Circle.Point.X = 8
	// equivalent to w.Circle.Point.Y = 8
	w.Radius = 5 // equivalent to w.Circle.Radius = 5
	w.Spokes = 20
	w.Circle.Point.X = 9 // We can still access the fields like this if we want to
	w.Circle.Point.Y = 9
	w.Circle.Radius = 7
	w.Spokes = 30

	// Defining a struct like this removes the ability to use the shorthand order-based literal syntax.
	// w = Wheel{8, 8, 5, 20}                       // compile error: unknown fields
	// w = Wheel{X: 8, Y: 8, Radius: 5, Spokes: 20} // compile error: unknown fields

	// Instead we must use the shape of the type declaration with one of these two forms:
	w = Wheel{Circle{Point{8, 8}, 5}, 20}
	w = Wheel{
		Circle: Circle{
			Point:  Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20, // As for all compsite literals, trailing commas are mandatory.
	}
	fmt.Printf("w: %#v\n", w)
}

// If we returned a value instead of a pointer,
// we would not be able to make assignments to returns.
func EmployeeByID(id int) *Employee {
	for _, e := range employees {
		if e.ID == id {
			return e
		}
	}
	return nil
}
