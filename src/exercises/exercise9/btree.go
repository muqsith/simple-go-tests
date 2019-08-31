package main

// alternative implementation : https://github.com/fgrehm/go-tour/blob/master/72-equivalent-binary-trees.go

import (
	"fmt"
	"time"

	"golang.org/x/tour/tree"
)

/*RecordKeeper blah blah*/
type RecordKeeper struct {
	r int
}

/*Walk blah blah*/
func Walk(t *tree.Tree, c chan int, rk RecordKeeper) {
	c <- t.Value
	rk.r--
	time.Sleep(time.Millisecond * 300)
	if t.Left != nil {
		rk.r++
		go Walk(t.Left, c, rk)
	}
	if t.Right != nil {
		rk.r++
		go Walk(t.Right, c, rk)
	}
	if rk.r == 0 {
		close(c)
	}
}

/*Same blah blah*/
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

	ar1Val, ar2Val := 0, 0
	valFound := false

	for i := 0; i < len(ar1); i++ {
		ar1Val = ar1[i]
		valFound = false
		for j := 0; j < len(ar2); j++ {
			ar2Val = ar2[j]
			if ar2Val == ar1Val && !valFound {
				valFound = true
				break
			}
		}
		isSame = isSame && valFound
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
