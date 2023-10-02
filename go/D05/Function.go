package main

import "fmt"

func main() {
	var x int32
	fmt.Printf("Enter two integers:")
	fmt.Scanf("%d", &x)
	fmt.Println(mx(10, 15))
	fmt.Println(mx(15, 10))
	fmt.Println(mxelement(12, 32, 10323, 655))

	///this is closure mean nested function. and recursive function
	/*
		func main() {
		  add := func(x, y int) int {
		    return x + y
		  }
		  fmt.Println(add(1,1))
		}

		func factorial(x uint) uint {
		  if x == 0 {
			return 1
		  }
		  return x * factorial(x-1)
		}
	*/
}
func mx(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
func mxelement(Arr ...int) int {
	x := Arr[0]
	for _, v := range Arr {
		if x < v {
			x = v
		}
	}
	return x
}
