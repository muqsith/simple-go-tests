package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type MyFloat float64

type Vertex struct {
	X, Y float64
}

type I interface {
	M()
}

func main() {
	var a Abser

	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f
	fmt.Printf("(%v, %T)\n", a, a)
	fmt.Println(a.Abs())

	a = &v
	fmt.Printf("(%v, %T)\n", a, a)
	fmt.Println(a.Abs())
	fmt.Println(v.Abs())

	var b *Vertex
	a = b
	fmt.Printf("(%v, %T)\n", a, a)
	fmt.Println(b.Abs())

	// NilInterfaceTest()

	EmptyInterfaceTest()
}

func (v *Vertex) Abs() float64 {
	if v == nil {
		fmt.Println("<nil>")
		return 0
	}
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func NilInterfaceTest() {
	var i I
	describe(i)
	i.M() // throws run time error
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func EmptyInterfaceTest() {
	var i interface{}

	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}
