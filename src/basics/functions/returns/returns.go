package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 5 / 10
	y = sum - x
	return
}

func returnTest(a, b int) (int, int) {
	return a + b, a - b
}

func main() {
	fmt.Println(split(200))
	x, y := returnTest(10, 5)
	fmt.Println(x, y)

	_, b := returnTest(3, 2)
	fmt.Println("b is ", b)
}
