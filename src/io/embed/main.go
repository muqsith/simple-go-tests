package main

import (
	"embed"
	"fmt"
	"log"
)

//go:embed tmp/data.txt
var f embed.FS

func main() {
	data, err := f.ReadFile("tmp/data.txt")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(string(data))
}
