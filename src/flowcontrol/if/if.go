package main

import (
	"fmt"
	"math"
)

func pow(x, y, limit float64) float64 {
	if v := math.Pow(x, y); v < limit {
		return v
	}
	return limit
}

func main() {
	fmt.Println(pow(2, 2, 10))
	fmt.Println(pow(3, 3, 20))
}
