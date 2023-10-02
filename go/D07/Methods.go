package main

import (
	"fmt"
	"math"
)

type Circle struct {
	r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}
func main() {
	c := new(Circle)
	c.r = 10
	fmt.Println(c.area())
	x := Circle{100}
	fmt.Println(x.area())
}
