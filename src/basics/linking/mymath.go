package main

import (
	"fmt"

	"muqsith/simple-go-tests/src/basics/linking/more"
)

/*

See below in order to compile this code

$ go build -o bin/links src/basics/linking/*.go

*/

func main() {
	fmt.Println("Mul(5, 4): ", Mul(5, 4))
	fmt.Println("Sub(2902,1029): ", Sub(2902, 1029))
	fmt.Println("Add(4, 8): ", more.Add(4, 8))
}
