package main

import "fmt"

func p(c chan int) {
	c <- 5
	c <- 10
	c <- 15
}

func main() {
	/*
		Buffered channel is like how much data in one shot we can write to and read from a channel
	*/
	c := make(chan int, 3)
	p(c)
	fmt.Println(<-c, <-c, <-c)
}
