package data

import "fmt"

var x int = 1

func GetX() int {
	return x
}

func SetX(newX int) {
	x = newX
}

// better
type Data struct {
	x int
}

func NewData(x int) *Data {
	return &Data{x: x}
}

func (d *Data) GetX() int {
	return d.x
}

func (d *Data) SetX(newX int) {
	d.x = newX
}

func (d *Data) String() string {
	return fmt.Sprintf("Data(x=%d)", d.x)
}
