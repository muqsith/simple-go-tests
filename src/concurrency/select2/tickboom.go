package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(300 * time.Millisecond)
	boom := time.After(2000 * time.Millisecond)

	for {
		select {
		case <-tick:
			{
				fmt.Println("tick.")
			}
		case <-boom:
			{
				fmt.Println("Boom!")
				return
			}
		default:
			{
				fmt.Println("    .")
				time.Sleep(50 * time.Millisecond)
			}
		}
	}
}
