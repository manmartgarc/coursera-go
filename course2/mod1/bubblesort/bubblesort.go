package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// BubbleSort sorts a slice of integers in ascending order in place.
func BubbleSort(slice []int) {
	sliceLength := len(slice)
	for i := 0; i < sliceLength; i++ {
		for j := sliceLength - 1; j > i; j-- {
			if slice[j] < slice[j-1] {
				Swap(slice, j-1)
			}
		}
	}
}

// swaps the elements at indices i and i+1 in a slice of integers.
func Swap(slice []int, i int) {
	slice[i], slice[i+1] = slice[i+1], slice[i]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	userSlice := make([]int, 0, 10)
	fmt.Print("Enter up to 10 integers separated by a space (or type x to exit): ")
	scanner.Scan()
	strInts := strings.Fields(scanner.Text())
	if len(strInts) > 10 {
		fmt.Println("Too many integers entered. Please enter up to 10 integers.")
		return
	}
	for _, strInt := range strInts {
		newInt, err := strconv.Atoi(strInt)
		if err != nil {
			if strings.ToLower(strInt) == "x" {
				fmt.Printf("Exiting...\n")
				break
			}
			fmt.Println("Invalid input. Please enter *only* integers.")
			return
		}
		userSlice = append(userSlice, newInt)
	}
	BubbleSort(userSlice)
	fmt.Printf("Sorted slice: %v\n", userSlice)
}
