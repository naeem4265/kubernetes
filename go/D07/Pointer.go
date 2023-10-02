package main

import "fmt"

func main() {
	ptr := new(int)
	*ptr = 10
	fmt.Println(*ptr)
	change(ptr)
	fmt.Println(*ptr)
}
func change(x *int) {
	*x = 1
}
