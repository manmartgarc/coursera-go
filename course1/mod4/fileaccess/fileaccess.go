package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// the class mentions ioutil package but it's deprecated since go 1.16
	// use os.ReadFile instead
	barr, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("File read: %s\n", string(barr))
	finalOut := append(barr, 'w', 'r', 'i', 't', 't', 'e', 'n', '\n')
	os.WriteFile("test2.txt", finalOut, 0644)

	// we can also use io package to read and write files but it's more complex
	file, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Printf("File read: %s", line)
	}

	f, err := os.Create("test3.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	writer := bufio.NewWriter(f)
	writer.WriteString("written\n")
	writer.Flush()
}
