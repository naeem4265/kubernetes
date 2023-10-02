package main

import (
	"fmt"
)

func main() {
	var x int64
	fmt.Scanf("%d", &x)
	fmt.Println("The number is")
	switch x {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("Unknown")
	}
}
