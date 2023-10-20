package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}
type rectangle struct {
	length, width, height float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (r rectangle) area() float64 {
	return r.length * r.width
}

func (c circle) volume() float64 {
	return math.Pi * c.radius * c.radius * c.radius * 4 / 3
}
func (r rectangle) volume() float64 {
	return r.length * r.width * r.height
}

type shape interface {
	area() float64
	volume() float64
}

func printShapeInfo(c shape) {
	fmt.Printf("Area of %T is %0.2f\n", c, c.area())
	fmt.Printf("Volume of %T is %0.2f\n", c, c.volume())
}

func main() {
	shapes := []shape{
		circle{10.0},
		rectangle{10, 20, 30},
		circle{20.0},
		rectangle{20, 30, 40},
	}
	for _, x := range shapes {
		printShapeInfo(x)
		fmt.Println()
	}
}
