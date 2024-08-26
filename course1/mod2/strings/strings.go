package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	// working with runes (unicode characters)
	fmt.Println(unicode.ToLower('a'))
	fmt.Println(unicode.ToUpper('a'))
	fmt.Println(string(unicode.ToUpper('a')))

	// working with strings
	if strings.Compare("Hello world!", "Hello, world!") == 0 {
		fmt.Println("Strings are equal")
	} else {
		fmt.Println("Strings are not equal")
	}

	if strings.Contains("Hello world!", "world") {
		fmt.Println("Contains")
	} else {
		fmt.Println("Does not contain")
	}
	fmt.Println(strings.Index("Hello, world!", "world"))

	// string manipulation (all return new strings)
	fmt.Println(strings.Replace("Hello, world!", "world", "gopher", 1))
	fmt.Println(strings.ToUpper("Hello, world!"))
	fmt.Println(strings.TrimSpace("  Hello, world!  "))

	// converting strings to/from things
	i, e := strconv.Atoi("10")
	if e != nil {
		panic(e)
	}
	fmt.Println(i)

	fmt.Println(strconv.FormatFloat(3.1415, 'e', 2, 32))
	f, _ := strconv.ParseFloat("3.1415", 32)
	fmt.Println(f)
}
