package main

import (
	"courserago/course2/mod3/encapsulation/data"
	"fmt"
)

func main() {
	fmt.Println(data.GetX())
	data.SetX(2)
	fmt.Println(data.GetX())

	d := data.NewData(3)
	fmt.Println(d.GetX())
	d.SetX(d.GetX() * 2)
	fmt.Println(d.GetX())
	fmt.Println(d)
}
