package main

import "fmt"

func f(ch chan int) {
	for i := 1; i < 10; i++ {
		ch <- i * i
	}
	close(ch)
}
func main() {
	ch := make(chan int)
	go f(ch)
	for {
		x, ok := <-ch
		if ok == true {
			fmt.Println("Data received : ", x)
		} else {
			fmt.Println("Break: ", x, ok)
			break
		}
	}
	fmt.Println("successful")
}
