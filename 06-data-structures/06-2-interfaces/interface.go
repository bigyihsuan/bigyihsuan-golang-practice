package main

import (
	"fmt"
	"math"
)

type Shape interface {
	SideCount() int
}

type Polygon interface {
	Perimeter() int
}

type Object interface {
	Area() float64
	Shape
	Polygon
}

// for Shape
func (p Pentagon) SideCount() int {
	return 5
}
// for Polygon
func (p Pentagon) Perimeter() int {
	return int(p) * p.SideCount()
}
// for Object
func (p Pentagon) Area() float64 {
	return 0.25 * math.Sqrt(5 * (5 + 2*math.Sqrt(5))) * float64(p)* float64(p)
}
type Pentagon int


func main() {
	pent := Pentagon(10)
	fmt.Printf("pent: %v\n", pent)
	fmt.Printf("pent.SideCount() (Shape): %v\n", pent.SideCount())
	fmt.Printf("pent.Perimeter() (Polygon): %v\n", pent.Perimeter())
	fmt.Printf("pent.Area(): %v\n", pent.Area())

	var shape Shape = Shape(pent)
	fmt.Printf("shape.SideCount(): %v\n", shape.SideCount())

	var polygon Polygon = Polygon(pent)
	fmt.Printf("polygon.Perimeter(): %v\n", polygon.Perimeter())

	var object Object = Object(pent)
	fmt.Printf("object.Area(): %v\n", object.Area())
	fmt.Printf("object.SideCount(): %v\n", object.SideCount())
	fmt.Printf("object.Perimeter(): %v\n", object.Perimeter())
}
