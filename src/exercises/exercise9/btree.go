package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, c chan int) {
	if t.Left != nil {
		go Walk(t.Left, c)
	}
	if t.Right != nil {
		go Walk(t.Right, c)
	}
	c <- t.Value
}

func main() {
	t := tree.New(1)
	c := make(chan int)

	go Walk(t, c)

	for i := range c {
		fmt.Println(i)
	}
}
