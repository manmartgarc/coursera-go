package main

import (
	"fmt"
	"math/rand"
)

type error interface {
	Error() string
}

type MyError struct {
	Msg string
}

func (e *MyError) Error() string {
	return e.Msg
}

func main() {
	var e error
	e = &MyError{"Something went wrong"}
	// we can use type assertion to check if the interface value is of a specific type
	// if the assertion fails, a panic will occur
	if me, ok := e.(*MyError); ok {
		fmt.Println(me.Msg)
	} else {
		fmt.Println("Not a MyError")
	}

	// we can do this with any boolean check, not just the type assertion
	if me, ok := e, true; ok && me.Error() == "Something went wrong" {
		fmt.Println("Something went wrong")
	}

	r := rand.Intn(10)

	if ok := r >= 5; ok {
		fmt.Println("Greater or equal to 5")
	} else {
		fmt.Println("Less than 5")
	}

	// we can also use a switch statement to check the type of the interface value
	switch err := e.(type) {
	case *MyError:
		fmt.Println(err.Msg)
	}
}
