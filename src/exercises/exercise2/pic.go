package main

import (
	"math"

	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	arr := make([][]uint8, dy, dy)
	for i := range arr {
		arr[i] = make([]uint8, dx, dx)
		for j := range arr[i] {
			// arr[i][j] = uint8(dx + dy)
			// arr[i][j] = uint8(dx * dy)
			arr[i][j] = uint8(math.Pow(float64(dx), float64(dy)))
		}
	}
	return arr
}

func main() {
	pic.Show(Pic)
	// remove "IMAGE:" and append "data:image/png;base64," to the output and see generated image in in browser
}
