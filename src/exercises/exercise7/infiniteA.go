package main

import (
	"fmt"

	"golang.org/x/tour/reader"
)

type InfiA byte

func (infiA InfiA) Read(b []byte) (int, error) {
	for i := 0; i < len(b); i += 1 {
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	b := make([]byte, 8)

	a := InfiA(0)

	for i := 0; i < 10; i += 1 {
		a.Read(b)
		fmt.Printf("%q\n", b)
	}

	reader.Validate(InfiA(0))
}
