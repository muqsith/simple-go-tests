package main

// courtesy : https://github.com/fgrehm/go-tour/blob/master/72-equivalent-binary-trees.go

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func walkTree(t *tree.Tree, c chan int) {
	if t.Left != nil {
		walkTree(t.Left, c)
	}
	c <- t.Value
	if t.Right != nil {
		walkTree(t.Right, c)
	}
}

func Walk(t *tree.Tree, c chan int) {
	walkTree(t, c)
	close(c)
}

func Same(t1, t2 *tree.Tree) bool {
	c1, c2 := make(chan int), make(chan int)
	go Walk(t1, c1)
	go Walk(t2, c2)

	for i := range c1 {
		if i != <-c2 {
			return false
		}
	}

	return true
}

func main() {
	t1 := tree.New(1)
	t2 := tree.New(1)
	t3 := tree.New(3)

	c1 := make(chan int)

	go Walk(t1, c1)

	for i := range c1 {
		fmt.Println(i)
	}

	fmt.Printf("(t1 == t2) = %v\n", Same(t1, t2))
	fmt.Printf("(t1 == t3) = %v\n", Same(t1, t3))
}
