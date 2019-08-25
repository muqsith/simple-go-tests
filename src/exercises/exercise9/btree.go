package main

// alternative implementation : https://github.com/fgrehm/go-tour/blob/master/72-equivalent-binary-trees.go

import (
	"fmt"
	"time"

	"golang.org/x/tour/tree"
)

type RecordKeeper struct {
	r int
}

func Walk(t *tree.Tree, c chan int, rk RecordKeeper) {
	c <- t.Value
	rk.r -= 1
	time.Sleep(time.Millisecond * 300)
	if t.Left != nil {
		rk.r += 1
		go Walk(t.Left, c, rk)
	}
	if t.Right != nil {
		rk.r += 1
		go Walk(t.Right, c, rk)
	}
	if rk.r == 0 {
		close(c)
	}
}

func Same(t1, t2 *tree.Tree) bool {
	isSame := true
	c1, c2 := make(chan int), make(chan int)
	rk1, rk2 := RecordKeeper{r: 1}, RecordKeeper{r: 1}
	go Walk(t1, c1, rk1)
	go Walk(t2, c2, rk2)

	ar1 := make([]int, 20)
	ar2 := make([]int, 20)

	for i := range c1 {
		ar1 = append(ar1, i)
	}

	for i := range c2 {
		ar2 = append(ar2, i)
	}

	ar1_val, ar2_val := 0, 0
	val_found := false

	for i := 0; i < len(ar1); i += 1 {
		ar1_val = ar1[i]
		val_found = false
		for j := 0; j < len(ar2); j += 1 {
			ar2_val = ar2[j]
			if ar2_val == ar1_val && !val_found {
				val_found = true
				break
			}
		}
		isSame = isSame && val_found
	}

	return isSame
}

func main() {
	t1 := tree.New(1)
	t2 := tree.New(1)
	t3 := tree.New(3)

	c1 := make(chan int)

	rk := RecordKeeper{r: 0}
	rk.r = 1
	go Walk(t1, c1, rk)

	for i := range c1 {
		fmt.Println(i)
	}

	fmt.Printf("(t1 == t2) = %v\n", Same(t1, t2))
	fmt.Printf("(t1 == t3) = %v\n", Same(t1, t3))
}
