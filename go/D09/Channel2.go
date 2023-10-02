package main

import "fmt"

func hello(x chan int) {
	fmt.Println("Hello")
	x <- 10
}

func main() {
	done := make(chan int)
	go hello(done)
	<-done
	fmt.Println("Main")
}
