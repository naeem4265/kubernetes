package main

import "fmt"

func main() {
	defer func() {
		str := "aneem"
		fmt.Println(str)
	}()
}
