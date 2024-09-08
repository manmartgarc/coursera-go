package main

import (
	"fmt"
	"math"
)

type Measurable interface {
	Area() float64
	Perimeter() float64
}

// no need to implement the interface explicitly, as long as the methods are implemented
// the compiler will automatically infer that the type satisfies the interface
type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r *Rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func measure(m Measurable) {
	fmt.Printf("Area: %.2f\n", m.Area())
	fmt.Printf("Perimeter: %2.f\n", m.Perimeter())
}

func main() {
	r := &Rectangle{Width: 3, Height: 4}
	c := &Circle{Radius: 5}

	fmt.Println("Rectangle:")
	measure(r)
	fmt.Println("Circle:")
	measure(c)
}
