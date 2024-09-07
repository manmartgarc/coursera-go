package main

import "fmt"

// lowercase function name means it's not exported
// uppercase function name means it's exported
// Exported functions can be used by other packages
func printHello() {
	fmt.Println("Hello, world!")
}

func main() {
	printHello()
}
