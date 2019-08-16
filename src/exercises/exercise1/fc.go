// exercise for flow control https://tour.golang.org/flowcontrol/8

package main

import (
	"fmt"
	"math"
)

/*
	A function that computes square root
*/
func Sqrt_Ten(x float64) float64 {
	z := 1.0

	for i := 0; i < 10; i += 1 {
		z -= ((z * z) - x) / (2 * z)
	}

	return z
}

func Sqrt(x float64, guess float64) (float64, int) {
	iters := 0
	z := guess
	z_previous := 0.0

	for z != z_previous {
		iters += 1
		z_previous = z
		z -= ((z * z) - x) / (2 * z)
	}

	return z, iters
}

func main() {
	num := float64(9)
	guess := num / 2
	result, iterations := Sqrt(num, guess)
	fmt.Printf("Square root of %v is %v\n", num, result)
	fmt.Println("Iterations: ", iterations)
	fmt.Printf("Actual square root of %v is %v\n", num, math.Sqrt(num))
}
