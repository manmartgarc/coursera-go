package main

import "fmt"

func main() {
	var x int32 = 1
	var y int64 = 2
	fmt.Println(int32(y) + x)
	fmt.Println((y) + int64(x))
}
