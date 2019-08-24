package main

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int) {
	var sum int
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {
	s := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31}

	c := make(chan int)

	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	time.Sleep(2000 * time.Millisecond)
	s1 := <-c
	time.Sleep(2000 * time.Millisecond)
	s2 := <-c

	var t int

	for _, v := range s {
		t += v
	}

	fmt.Printf("Total(t): %d\n", t)

	fmt.Printf("Total(s1+s2): %d\n", s1+s2)

}
