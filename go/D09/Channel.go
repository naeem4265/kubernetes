package main

import "fmt"

func main() {
	a := make(chan int)
	fmt.Printf("Type of chan %T", a)
}
