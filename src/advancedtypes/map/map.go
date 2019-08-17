package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["kurnool"] = Vertex{
		40.68433,
		-74.39967,
	}
	fmt.Println(m["kurnool"])

	mapLiteralsTest()
	mutatingMaps()
}

func mapLiteralsTest() {
	// map keys could be anything not just strings
	m := map[int]string{
		1: "a",
		2: "b",
	}

	fmt.Println(m[2])
}

func mutatingMaps() {
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2

	delete(m, "b")

	v, ok := m["b"]
	fmt.Println(v, ok)
	v, ok = m["a"]
	fmt.Println(v, ok)
}
