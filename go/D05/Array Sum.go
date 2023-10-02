package main

import "fmt"

func main() {
	///
	Array1 := [5]int{10, 20, 30, 40, 50}
	for i := 0; i < 5; i++ {
		fmt.Println(i, Array1[i])
	}
	var Arr [5]int
	Arr[2] = 100
	for _, value := range Arr {
		fmt.Println(value)
	}

	/// Slice
	A := make([]int, 2) // Make a Slice with length 2
	for _, x := range A {
		fmt.Println(x)
	}
	B := Array1[1:3] // Make a Slice with element Array index 1 to Array index 3
	for _, x := range B {
		fmt.Println(x)
	}
	C := append(B, 5) // Insert a Slice to another Slice and one element 5
	for _, x := range C {
		fmt.Println(x)
	}
	copy(A, C) // Copy Slice C to A
	for _, x := range A {
		fmt.Println(x)
	}
	/// Mapping
	mp := make(map[string]int) // Declare mp map to point integer instead of string
	mp["Naeem"] = 10050
	fmt.Println(mp["Naeem"])
}
