package main

import "fmt"

func main() {
	var Array [10]int64
	Array[0] = 5
	Array[5] = 6
	for i := 0; i < 10; i++ {
		f(Array[i])
	}
}
func f(x int64) {
	if x%2 == 0 {
		fmt.Println(x, " is even")
	} else {
		fmt.Println(x, "is odd")
	}
}
