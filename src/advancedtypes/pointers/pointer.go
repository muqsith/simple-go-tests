package main

import (
	"fmt"
	"math"
)

/*Circle is just a circle*/
type Circle struct {
	radius float64
	area   float64
}

func (c *Circle) calculateArea() {
	c.area = math.Pi * c.radius * c.radius
}

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

	c1 := Circle{
		radius: 7.0,
	}

	pc1 := &c1
	pc2 := &c1

	c1.calculateArea()

	fmt.Println(pc1.area)

	fmt.Println(pc1 == pc2)

	var pa *int

	pa = &j

	fmt.Println("adress pa referring to: ", pa)
}
