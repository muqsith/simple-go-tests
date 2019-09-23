package main

import (
	"fmt"
	"strings"
)

func main() {
	var a [2]string
	a[0] = "muqsith"
	fmt.Println(a[0])
	fmt.Println(a[1])

	var primes = [10]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	fmt.Println(primes)

	composites := [5]int{1, 2, 4, 6, 8}
	fmt.Println(composites)

	// slices
	someprimes := primes[2:5]
	fmt.Println(someprimes)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	b := names[1:3]
	fmt.Println(b)
	b[0] = "muqsith"
	fmt.Println(b)

	p := [...]string{"x", "y"}
	fmt.Println("p: ", p)

	q := []string{"a", "b", "c", "d"}
	fmt.Println(q)

	// full slice
	r := q[:]
	fmt.Println(r)

	const Enone int = 1
	const Eio = 2
	const Einval = 4

	alienArray := [...]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}

	for key, val := range alienArray {
		fmt.Println(key, val)
	}

	alterSlices()
	nilTest()
	makeTest()
	slicesOfSlices()
	appendSlices()
	copySlices()

	var test []interface{}
	fmt.Println("test len: ", len(test))
}

func alterSlices() {
	fmt.Printf("\n\nalter slices test\n")
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice("s", s)

	s = s[:0]
	printSlice("s[:0]", s)
	fmt.Println("s[3:4]", s[3:4])

	s = s[:4]
	printSlice("s[:4]", s)

	s = s[2:]
	printSlice("s[2:]", s)
}

func printSlice(text string, s []int) {
	fmt.Printf("length of %s is %d, capacity of %s is %d\n", text, len(s), text, cap(s))
}

func nilTest() {
	fmt.Printf("\n\nnil test\n")
	var s []int
	fmt.Printf("len of s is %d, cap of s is %d\n", len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

func makeTest() {
	fmt.Printf("\n\nmake test\n")
	s := make([]int, 0, 5)
	printSlice("s", s)

	s1 := s[:2]
	printSlice("s1", s1)

	s2 := s1[1:4]
	printSlice("s2", s2)
}

// array of arrays
func slicesOfSlices() {
	fmt.Printf("\n\nslices of slices test\n")
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

}

func appendSlices() {
	fmt.Printf("\n\nappend slices test\n")

	var s []int
	printSlice("s", s)

	s = append(s, 0)
	printSlice("s", s)

	s = append(s, 1)
	printSlice("s", s)

	s = append(s, 2, 3, 4, 5)
	printSlice("s", s)

	fmt.Println(s)
}

func copySlices() {
	fmt.Printf("\n\ncopy slices test\n")

	s1 := []int{0, 2, 4, 6, 8, 10}
	s2 := []int{1, 3, 5, 7, 9, 11}

	s := make([]int, len(s1)+len(s2))

	copy(s, s1)
	copy(s[len(s1):], s2)

	for _, i := range s {
		fmt.Printf("%d\n", i)
	}
}
