package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func a() {
	v1 := Vertex{1, 2}
	v1.X = 4
	fmt.Println("v1.X", v1.X)
}

func b() {
	v2 := Vertex{4, 5}
	p := &v2
	fmt.Printf("Type of p is %T\n", p)
	// observe below the way of accessing v2 properties via pointer
	(*p).X = 1e9
	p.Y = p.Y * 300
	fmt.Println("v2: ", v2)
}

func c() {
	var (
		v1 = Vertex{1, 2}
		v2 = Vertex{X: 10, Y: 20}
		v3 = Vertex{}
		p  = &Vertex{35, 45}
	)
	fmt.Println(v1, v2, v3, p)
}

func main() {
	fmt.Println(Vertex{2, 3})
	a()
	b()
	c()
}
