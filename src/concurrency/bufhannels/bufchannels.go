package main

import (
	"fmt"
	"time"
)

func sender(c chan<- int) {
	for i := 0; i < 10; i++ {

		v1 := i + 5
		fmt.Println("pushed " + fmt.Sprint(v1) + " to the channel.")
		c <- v1

		v2 := i + 10
		fmt.Println("pushed " + fmt.Sprint(v2) + " to the channel.")
		c <- v2

		v3 := i + 15
		fmt.Println("pushed " + fmt.Sprint(v3) + " to the channel.")
		c <- v3

	}
	close(c)
}

func main() {
	/*
		https://go.dev/tour/concurrency/3
		Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.

		Own words:
		Sends will not be blocked till the buffer reaches it's capacity and
		recieves will only start when then buffer has reached it's capacity.
	*/

	// remove 3 in below line to observe difference in behavior
	c := make(chan int, 3)
	go sender(c)

	for v := range c {
		time.Sleep(time.Second * 2)
		fmt.Println("Received: ", v)
	}

}
