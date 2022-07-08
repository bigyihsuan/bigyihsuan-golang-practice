package main

import (
	"fmt"
	"math"
)

type Solid interface {
	Volume() float64
}

type Cube struct {
	length float64
}
type Box struct {
	length float64
	width  float64
	height float64
}
type Sphere struct {
	radius float64
}

func (c Cube) Volume() float64 {
	return c.length * c.length * c.length
}
func (b Box) Volume() float64 {
	return b.length * b.width * b.height
}
func (s Sphere) Volume() float64 {
	return float64(4) / float64(3) * math.Pi * s.radius * s.radius * s.radius
}

func main() {
	c := Cube{length: 5}
	b := Box{length: 5, width: 3, height: 4}
	s := Sphere{radius: 5}

	fmt.Printf("c.Volume(): %v\n", c.Volume())
	fmt.Printf("b.Volume(): %v\n", b.Volume())
	fmt.Printf("s.Volume(): %v\n", s.Volume())
}
