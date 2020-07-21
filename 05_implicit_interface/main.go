// Description:
//	- Demo implicit interface
//
// Output:
//	area of *main.Rect = 12
//	area of *main.Circle = 78.53981633974483
package main

import (
	"fmt"
	"math"
)

// Shape interface
type Shape interface {
	area() float64
}

// Rect rectangle struct
type Rect struct {
	width  float64
	height float64
}

// Circle struct
type Circle struct {
	radius float64
}

// pass by reference for demo purpose
func (r *Rect) area() float64 {
	return r.width * r.height
}

// pass by reference for demo purpose
func (c *Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	r := Rect{3, 4}
	c := Circle{5}
	// use address of (&) as pointer receiver
	shapes := []Shape{&r, &c}

	for _, shape := range shapes {
		fmt.Printf("area of %T = %v \n", shape, shape.area())
	}
}
