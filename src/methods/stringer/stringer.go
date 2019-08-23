package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Abdul Muqsith", 33}
	z := Person{"Mohammed Irfan", 32}
	fmt.Println(a, z)
}
