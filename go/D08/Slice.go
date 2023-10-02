package main

import "fmt"

func main() {
	arr := []string{"a", "b", "c", "d", "e"}
	fmt.Println(arr)
	arr2 := arr[:3]
	fmt.Println(arr2)
	arr2 = append(arr2, "x", "y", "z")
	fmt.Println(arr2)
	arr3 := make([]string, 2, 5)
	fmt.Println(arr3)
	fmt.Println(len(arr3), cap(arr3))
}
