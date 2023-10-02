package main

import "fmt"

type student struct {
	name string
	id   int
	cgpa float64
}

func (x *student) Print() {
	fmt.Println(x.name)
}

type duet struct {
	student
	university string
}

func (x *duet) Out() {
}

func main() {
	x := new(duet)
	x.student.name = "naeem"
	x.id = 174078
	x.cgpa = 3.27
	x.university = "DUET"
	x.Print()
}
