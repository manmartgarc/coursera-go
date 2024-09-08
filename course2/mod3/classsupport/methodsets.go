package main

import (
	"fmt"
	"math"
)

type Person struct {
	FirstName string
	LastName  string
}

// we can associate methods with a type, this is called a method set of the type,
// and yes, we are inching towards classes and interfaces.
func (p Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

func (p Person) Greeting() string {
	return "Hello, " + p.FullName()
}

type Point struct {
	x, y float64
}

// here we make use of the pointer reference to the Point object, so that we
// can modify the object in place, rather than creating a copy of it.
func (p *Point) DistToOrig() float64 {
	t := math.Pow(p.x, 2) + math.Pow(p.y, 2)
	return math.Sqrt(t)
}

func main() {
	p := Person{"John", "Doe"}
	// the one gotcha here is that the implicit parameter is passed to the
	// method by value; so it actually is a copy of the original object.
	// to avoid this, we need to define our method set with pointer references.
	fmt.Println(p.FullName())
	fmt.Println(p.Greeting())

	point := Point{3, 8}
	fmt.Println(point.DistToOrig())
}
