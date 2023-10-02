package main

import "fmt"

func main() {
	x := 1
	y := 2
	fmt.Println(x, y)
	Swap(&x, &y)
	fmt.Println(x, y)
}
func Swap(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
}
