package main

import (
	"fmt"
	"math"
)

type Circle struct {
	r float64
}

func main() {
	c := new(Circle)
	c.r = 10
	fmt.Println(area(c))
	var x Circle
	x.r = 100
	fmt.Println(area(&x))
}
func area(c *Circle) float64 {
	return math.Pi * c.r * c.r
}
