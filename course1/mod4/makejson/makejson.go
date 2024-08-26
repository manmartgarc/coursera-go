package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// Write a program which prompts the user to first enter a name, and then
// enter an address. Your program should create a map and add the name and
// address to the map using the keys “name” and “address”, respectively.
// Your program should use Marshal() to create a JSON object from the map,
// and then your program should print the JSON object.
func main() {
	jsonMap := make(map[string]string)

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter a name: ")
	scanner.Scan()
	jsonMap["name"] = scanner.Text()
	fmt.Print("Enter an address: ")
	scanner.Scan()
	jsonMap["address"] = scanner.Text()

	barr, err := json.Marshal(jsonMap)
	if err != nil {
		panic(err)
	}
	fmt.Printf("JSON object: %s\n", string(barr))
}
