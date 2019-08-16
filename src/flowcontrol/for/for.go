package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i <= 100; i += 1 {
		sum += i
	}
	fmt.Println(sum)

	sum2 := 1
	for sum2 < 10000 {
		sum2 += sum2
	}
	fmt.Println(sum2)

	// forever

	sum3 := 1
	for {
		sum3 += sum3
		if sum3 > 5000 {
			break
		}
	}
	fmt.Println(sum3)
}
