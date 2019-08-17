package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	prev, curr, next := 0, 0, 1
	return func() int {
		prev = curr
		curr = next
		next = prev + curr
		return prev
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
