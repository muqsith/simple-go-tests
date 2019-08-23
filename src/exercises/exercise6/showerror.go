package main

import (
	"fmt"
	"math"
)

type Sqrt float64

func (e *Sqrt) Error() string {
	return fmt.Sprintf("Cannot square root negative number: %v ", float64(*e))
}

func (s Sqrt) Sqrt() (float64, error) {
	if s < 0 {
		return 0, &s
	}
	return math.Sqrt(float64(s)), nil
}

func main() {

	n := Sqrt(-4.0)

	root, err := n.Sqrt()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("root of %v is %v\n", n, root)
	}
}
