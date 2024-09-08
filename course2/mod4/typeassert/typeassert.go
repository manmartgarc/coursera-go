package main

import (
	"fmt"
	"math"
)

type Measurable interface {
	area() float64
}

type Rectangle struct {
	width, height float64
}

func (r *Rectangle) area() float64 {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func (c *Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func main() {
	var m Measurable
	// m = &Rectangle{3, 4}
	m = &Circle{5}
	// we can use type assertion to check if the interface value is of a specific type
	// if the assertion fails, a panic will occur
	if c, ok := m.(*Circle); ok {
		fmt.Printf("Circle: %.2f\n", c.area())
	} else {
		fmt.Println("Not a Circle")
	}

	// we can also use a switch statement to check the type of the interface value
	switch sh := m.(type) {
	case *Rectangle:
		fmt.Printf("Rectangle: %.2f\n", sh.area())
	case *Circle:
		fmt.Printf("Circle: %.2f\n", sh.area())
	}
}
