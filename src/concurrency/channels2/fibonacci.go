package main

import (
	"fmt"
	"time"
)

func fib(n int, c chan int) {
	prev, curr, next := 0, 0, 1
	for i := 0; i < n; i += 1 {
		prev = curr
		c <- curr
		curr = next
		next = prev + curr
	}
	close(c)
}

func main() {
	c := make(chan int, 50)
	go fib(cap(c), c)
	for i := range c {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(i)
	}
}
