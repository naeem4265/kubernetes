package main

import "fmt"

func main() {
	A := [10]int{
		10, 20, 4, 235, 3,
		225, 256, 1, 32, 45,
	}
	var mn int
	mn = A[0]
	for _, x := range A {
		if mn > x {
			mn = x
		}
	}
	fmt.Println(mn)
}
