package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

var digitRegexp = regexp.MustCompile("[0-9]+")

func FindDigits(filename string) []byte {
	b, _ := ioutil.ReadFile(filename)
	return digitRegexp.Find(b)
}

func main() {
	filename := "sample.txt"
	digits := FindDigits(filename)
	if digits != nil {
		fmt.Println("Found digits:", string(digits))
	} else {
		fmt.Println("No digits found.")
	}
}
