package main

import "fmt"

type Person struct {
	name string
	age  int
}

func newPerson(name string, age int) *Person {
	p := Person{name: name, age: age}
	return &p
}

func main() {
	// can be defined via literal
	p1 := Person{
		name: "John",
		age:  30,
	}
	// or via new
	p2 := new(Person)
	p2.name = "Jane"
	p2.age = 25
	// or via constructor
	p3 := newPerson("Jack", 35)
	fmt.Printf("p1: %v\n", p1)
	fmt.Printf("p2: %v\n", *p2)
	fmt.Printf("p3: %v\n", *p3)
}
