package main

import (
	"fmt"
	"time"
)

func printDay() int {
	fmt.Println(time.Now().Weekday())
	return time.Now().Year()
}

func a() {
	for i := 0; i < 10; i += 1 {
		defer fmt.Println(i)
	}
}

func b(j int) (i int) {
	i = j
	// the below defer function catches the returning 'i' and increments it
	defer func() {
		i += 1
	}()
	return
}

func main() {
	defer fmt.Println(printDay())

	fmt.Println("Hello World!")

	a()
	fmt.Println("b(5)", b(5))
}
