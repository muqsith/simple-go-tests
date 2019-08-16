package main

import "fmt"

var c, python, java bool

var p, q int = 1, 2

func main() {
	var i int
	fmt.Println(i, c, python, java)

	var l, fruit = 10, "Mango"
	fmt.Println(p, q, l, fruit)

	// implicit type - only available inside functions
	k := 32
	fmt.Println(k)
}
