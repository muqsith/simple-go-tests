package main

import (
	"context"
	"fmt"
	"time"
)

const shortDuration = 5 * time.Second

func main() {
	// Pass a context with a timeout to tell a blocking function that it
	// should abandon its work after the timeout elapses.
	ctx, cancel := context.WithTimeout(context.Background(), shortDuration)
	defer cancel()

	select {
	case <-time.After(2 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println("context done: ", ctx.Err()) // prints "context deadline exceeded"
	}

}
