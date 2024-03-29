package main

import (
	"context"
	"fmt"
	"time"
)

func channelReceiver(c <-chan int) {
	for {
		time.Sleep(500 * time.Millisecond)
		v := <-c
		fmt.Println("channelReceiver received: ", v)
	}
}

func main() {
	// gen generates integers in a separate goroutine and
	// sends them to the returned channel.
	// The callers of gen need to cancel the context once
	// they are done consuming generated integers not to leak
	// the internal goroutine started by gen.
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		go channelReceiver(dst)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return // returning not to leak the goroutine
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers

	for n := range gen(ctx) {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(n)
		if n > 5 {
			break
		}
	}
}
