package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Write a program which prompts the user to enter a string. The program
// searches through the entered string for the characters ‘i’, ‘a’, and ‘n’.
// The program should print “Found!” if the entered string starts with the
// character ‘i’, ends with the character ‘n’, and contains the character ‘a’.
// The program should print “Not Found!” otherwise. The program should not be
// case-sensitive, so it does not matter if the characters are upper-case or
// lower-case.
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter a string: ")
	scanner.Scan()
	userStr := strings.ToLower(strings.TrimSpace(scanner.Text()))
	if strings.HasPrefix(userStr, "i") && strings.HasSuffix(userStr, "n") && strings.Contains(userStr, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
