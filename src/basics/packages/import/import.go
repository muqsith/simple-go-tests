package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// rand.Seed(42)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		fmt.Println("My favorite number is", rand.Intn(500))
	}
}
