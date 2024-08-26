package main

import "fmt"

func main() {
	// can be defined via literal
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	// or via make
	m2 := make(map[string]int)
	m2["one"] = 1
	m2["two"] = 2
	m2["three"] = 3
	fmt.Printf("value of m[\"one\"]: %d\n", m["one"])
	fmt.Printf("value of m2[\"one\"]: %d\n", m2["one"])

	// we can delete a key
	delete(m, "one")
	fmt.Printf("value of m[\"one\"] after delete: %d\n", m["one"])

	// two value assignment
	id, p := m2["one"]
	if p {
		fmt.Printf("id: %d, p: %t\n", id, p)
	} else {
		fmt.Printf("key not found\n")
	}

	// iterating over a map
	for key, value := range m2 {
		fmt.Printf("key: %s, value: %d\n", key, value)
	}
}
