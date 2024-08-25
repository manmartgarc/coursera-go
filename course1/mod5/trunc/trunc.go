package main

import "fmt"

// Prompts the user to enter a floating point number and prints the integer
// which is a truncated version of the floating point number that was entered.
func main() {
	var num float64
	fmt.Print("Enter a floating point number: ")
	nRead, err := fmt.Scan(&num)
	if nRead != 1 || err != nil {
		fmt.Println("Invalid input. Please enter a floating point number.")
	} else {
		fmt.Printf("Truncated number: %d\n", int(num))
	}
}
