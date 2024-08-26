package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

// Write a program which prompts the user to enter integers and stores the
// integers in a sorted slice. The program should be written as a loop. Before
// entering the loop, the program should create an empty integer slice of size
// (length) 3. During each pass through the loop, the program prompts the user
// to enter an integer to be added to the slice. The program adds the integer
// to the slice, sorts the slice, and prints the contents of the slice in
// sorted order. The slice must grow in size to accommodate any number of
// integers which the user decides to enter. The program should only quit
// (exiting the loop) when the user enters the character ‘X’ instead of an
// integer.
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	userSlice := make([]int, 0, 3)
	enter := "Enter an integer (or type x to exit): "
	for fmt.Print(enter); scanner.Scan(); fmt.Print(enter) {
		newInt, err := strconv.Atoi(scanner.Text())
		if err != nil {
			if strings.ToLower(scanner.Text()) == "x" {
				fmt.Printf("Exiting...\n")
				break
			}
			fmt.Println("Invalid input. Please enter an integer.")
			continue
		}
		userSlice = append(userSlice, newInt)
		slices.Sort(userSlice)
		fmt.Printf("Sorted slice: %v\n", userSlice)
	}
}
