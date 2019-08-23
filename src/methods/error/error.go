package main

import "fmt"

type Circle struct {
	radius int
	area   float64
}

func (c *Circle) Error() string {
	return fmt.Sprintf("Raidus: %v", c.radius)
}

func main() {
	c := &Circle{
		radius: 7,
	}

	fmt.Println(c)
}
