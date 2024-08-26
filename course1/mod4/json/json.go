package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name      string
	Age       int
	candidate bool
}

func NewPerson(name string, age int) *Person {
	return &Person{name, age, false}
}

func main() {
	p := NewPerson("John", 30)
	p.candidate = true
	barr, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	fmt.Println("Note that the candidate field is not exported")
	fmt.Printf("Person: %v\n", *p)
	fmt.Printf("Person: %s\n", string(barr))

	var p2 Person
	err = json.Unmarshal(barr, &p2)
	if err != nil {
		panic(err)
	}
	fmt.Println("After unmarshalling")
	fmt.Printf("Person: %v\n", p2)
}
