package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i += 1 {
		time.Sleep(300 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}
