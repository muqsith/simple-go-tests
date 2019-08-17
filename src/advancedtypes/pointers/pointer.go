package main

import "fmt"

func main() {
	var i int = 42
	var p *int = &i
	fmt.Println("Address of i: ", p)
	fmt.Println("Value at ", p, " is ", *p)
	fmt.Printf("Type of p is %T \n", p)

	j := 36
	p = &j
	fmt.Println("Value at ", p, " is ", *p)
	*p = *p + 45
	fmt.Println("Value at ", p, " is ", *p)
}
