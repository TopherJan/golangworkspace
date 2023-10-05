package main

import (
	"fmt"
)

// triangle represents a triangle object
type triangle struct {
	height float64
	base   float64
}

// square represents a square object
type square struct {
	sideLength float64
}

// shape represents an interface for all the shapes
type shape interface {
	getArea() float64
}

func main() {
	t := triangle{
		base:   10,
		height: 20,
	}

	s := square{
		sideLength: 5,
	}

	printArea(t)
	printArea(s)
}

// getArea returns the area of the triangle
func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

// getArea returns the area of the square
func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

// printArea prints out the area of the shape
func printArea(s shape) {
	fmt.Printf("Area of %T: %f\n", s, s.getArea())
}
