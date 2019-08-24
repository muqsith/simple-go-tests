package main

import (
	"fmt"
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
}

func (pic Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 100, 100)
}

func (pic Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (pic Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x), uint8(y), 255, 255}
}

func main() {
	fmt.Println()
	m := Image{}
	pic.ShowImage(m)

	/*
		"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAGQAAABkCAIAAAD/gAIDAAAA7UlEQVR4nOzQQQnAQAADwRTOv+WWOrj9zxAiYM/2Ppvd7PzHHbECsQKxArECsQKxArECsQKxArECsQKxArECsQKxArECsQKxArECsQKxArECsQKxArECsQKxArECsQKxArECsQKxArECsQKxArECsQKxArECsQKxArECsQKxArECsQKxArECsQKxArECsQKxArECsQKxArECsQKxArECsQKxArECsQKxArECsQKxArECsQKxArECsQKxArECsQKxArECsQKxArECsQKxArECsQKxArECsQKxArECsQKxArECsQKxArECsYIvAAD//2FEA1bTHkkPAAAAAElFTkSuQmCC"
	*/
}
