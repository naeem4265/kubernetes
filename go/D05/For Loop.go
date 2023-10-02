package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	var Arr [5]int
	Arr[2] = 100
	for _, value := range Arr {
		fmt.Println(value)
	}
}
