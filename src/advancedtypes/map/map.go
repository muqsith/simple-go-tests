package main

import "fmt"

/*Vertex is a vertex*/
type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func updateMap(m map[string]string, key, value string) {
	m[key] = value
}

func main() {
	m = make(map[string]Vertex)
	m["kurnool"] = Vertex{
		40.68433,
		-74.39967,
	}
	fmt.Println(m["kurnool"])

	mapLiteralsTest()
	mutatingMaps()

	// maps DO NOT copy when assigned, maps are always references
	// see https://dave.cheney.net/2017/04/30/if-a-map-isnt-a-reference-variable-what-is-it
	place := make(map[string]string)
	placeCopy := place

	place["name"] = "Paris"
	updateMap(place, "country", "France")

	fmt.Println(place)
	fmt.Println(placeCopy)

	var place2 = map[string]string{
		"name": "Berlin",
	}

	updateMap(place2, "country", "Germany")

	fmt.Println(place2)

	var nilMap map[string]interface{}
	for k, v := range nilMap {
		fmt.Println("k: ", k, ", v: ", v)
	}
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
