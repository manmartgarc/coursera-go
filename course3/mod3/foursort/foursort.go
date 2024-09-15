package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"
)

func sort(xs []int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Sorting", xs)
	slices.Sort(xs)
}

func chunkSlice(s []int) [][]int {
	chunkSize := len(s) / 4
	chunks := make([][]int, 4)
	for i := 0; i < 4; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == 3 {
			end = len(s)
		}
		chunks[i] = s[start:end]
	}
	return chunks
}

func merge(chunks [][]int) []int {
	result := make([]int, 0)
	for len(chunks[0]) > 0 || len(chunks[1]) > 0 || len(chunks[2]) > 0 || len(chunks[3]) > 0 {
		min := 0
		for i := 1; i < 4; i++ {
			if len(chunks[i]) > 0 && (len(chunks[min]) == 0 || chunks[i][0] < chunks[min][0]) {
				min = i
			}
		}
		result = append(result, chunks[min][0])
		chunks[min] = chunks[min][1:]
	}
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Type a list of space-separated integers to sort. (type 'q' to quit)")
	scanner.Scan()
	strInts := strings.Fields(scanner.Text())
	if len(strInts) == 0 || strInts[0] == "q" {
		return
	}
	s := make([]int, 0, len(strInts))

	for _, strInt := range strInts {
		newInt, err := strconv.Atoi(strInt)
		if err != nil {
			fmt.Println("Invalid input. Please enter *only* integers.")
			return
		}
		s = append(s, newInt)
	}
	wg := &sync.WaitGroup{}
	wg.Add(4)
	chunks := chunkSlice(s)
	for chunk := range chunks {
		go sort(chunks[chunk], wg)
	}
	wg.Wait()
	fmt.Printf("Sorted slice: %v\n", merge(chunks))
}
