package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

type Pint int

func (p Pint) Flint() Pint {
	return p + 2
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	v.Scale(3)
	fmt.Println(v.Abs())

	p := Pint(2)
	fmt.Println(p.Flint())
}
