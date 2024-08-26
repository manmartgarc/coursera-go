package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Person struct {
	fname string
	lname string
}

func (p Person) String() string {
	return fmt.Sprintf("Person(fname: %s, lname: %s)", p.fname, p.lname)
}

func newPerson(fname, lname string) Person {
	if len([]rune(fname)) > 20 {
		fname = fname[:20]
	}
	if len([]rune(lname)) > 20 {
		lname = lname[:20]
	}
	return Person{fname: fname, lname: lname}
}

// Write a program which reads information from a file and represents it
// in a slice of structs. Assume that there is a text file which contains a
// series of names. Each line of the text file has a first name and a last
// name, in that order, separated by a single space on the line.

// Your program will define a name struct which has two fields, fname for
// the first name, and lname for the last name. Each field will be a string
// of size 20 (characters).

// Your program should prompt the user for the name of the text file. Your
// program will successively read each line of the text file and create a
// struct which contains the first and last names found in the file. Each
// struct created will be added to a slice, and after all lines have been
// read from the file, your program will have a slice containing one struct
// for each line in the file. After reading all lines from the file, your
// program should iterate through your slice of structs and print the first
// and last names found in each struct.
func main() {
	fmt.Print("Enter the name of the file: ")
	var filename string
	_, err := fmt.Scan(&filename)
	if err != nil {
		log.Fatalf("failed to read filename: %v", err)
	}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	persons := make([]Person, 0, 25)
	for scanner.Scan() {
		p := newPerson("", "")
		_, err := fmt.Sscanf(scanner.Text(), "%s %s", &p.fname, &p.lname)
		if err != nil {
			log.Fatalf("failed to read name: %v", err)
		}
		persons = append(persons, p)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("failed to scan file: %v", err)
	}

	for _, p := range persons {
		fmt.Println(p)
	}

}
