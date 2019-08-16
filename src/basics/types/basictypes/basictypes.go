package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool   = true
	MaxInt uint64 = 1<<64 - 1
	z             = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type %T, Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type %T, Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type %T, Value: %v\n", z, z)

	// default values
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("i = %v, f = %v, b = %v, s = %q\n", i, f, b, s)
}
