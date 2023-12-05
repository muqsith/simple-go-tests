package main

import (
	"fmt"
)

func processReceiver(c chan int) {
	for {
		v := <-c
		fmt.Println("processReceiver received: ", v)
	}
}

func main() {
	/*
		https://go.dev/tour/concurrency/3
		Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.
	*/

	c := make(chan int, 3)

	go processReceiver(c)

	for i := 0; i < 10; i++ {
		fmt.Println("main sending: ", i)
		c <- i
	}

}
