package main

import "fmt"

func main() {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	c := make(chan int)
	i := 0
	var v []int
	go g(0, c)
	for len(v) < 3 {
		v = append(v, <-c)
		i++
	}
	fmt.Println("Returned normally from g.")
}

func g(i int, c chan int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
			return
		}
	}()
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
		//return
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i+1, c)
	c <- 1
}
