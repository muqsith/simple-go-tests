package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

func main() {
	type Inventory struct {
		Material string
		Count    uint
	}

	sweaters := Inventory{"wool", 17}
	tmpl, err := template.New("wool").Parse("{{.Count}} items are made of {{.Material}}\n")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		panic(err)
	}

	cottonShirts := Inventory{"cotton", 45}
	tmpl2, err := template.New("cotton").Parse(`We have {{.Count}} {{.Material}} shirts`)
	var b strings.Builder
	err = tmpl2.Execute(&b, cottonShirts)
	if err != nil {
		panic(err)
	}
	fmt.Println(b.String())
}
