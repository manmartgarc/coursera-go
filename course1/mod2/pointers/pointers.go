package main

import "fmt"

func main() {
	var x int = 1
	var y int
	var ip *int

	fmt.Println("Value of x is: ", x)
	fmt.Println("Value of y is: ", y)
	fmt.Println("Value of *ip is: <nil>")
	fmt.Println("Value of ip is: ", ip)

	ip = &x
	y = *ip
	fmt.Println("Value of x is: ", x)

	fmt.Println("Value of x is: ", x)
	fmt.Println("Value of y is: ", y)
	fmt.Println("Value of *ip is: ", *ip)
	fmt.Println("Value of ip is: ", ip)

	// we can also use new keyword to create a pointer
	ptr := new(int)
	*ptr = 3
	fmt.Printf("Pointer at address %p points to value %d\n", ptr, &*ptr)
}
