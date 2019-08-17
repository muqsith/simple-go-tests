package main

import "fmt"

var squares = []int{1, 4, 9, 16, 25, 36, 49, 64, 81, 100, 121}

func main() {
	for i, v := range squares {
		fmt.Printf("value at %d is %d\n", i, v)
	}

	for _, v := range squares {
		fmt.Printf("v = %d\n", v)
	}

	for i := range squares {
		fmt.Printf("index %d\n", i)
	}
}
