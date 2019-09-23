package main

import (
	"fmt"
	"math/cmplx"
)

var (
	toBe   bool   = true
	maxInt uint64 = 1<<64 - 1
	z             = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type %T, Value: %v\n", toBe, toBe)
	fmt.Printf("Type %T, Value: %v\n", maxInt, maxInt)
	fmt.Printf("Type %T, Value: %v\n", z, z)

	// default values
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("i = %v, f = %v, b = %v, s = %q\n", i, f, b, s)

	s1 := "hello"
	s2 := "world"
	s3 := s1 + ", " + s2

	fmt.Println(s3)

	s4 := "hello worl\n"
	fmt.Printf("%q", s4)
	fmt.Println()
	// string equality

	n1 := "muqsith"
	n2 := "muqsith"
	fmt.Println("string equality: ", n1 == n2)
}
