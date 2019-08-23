package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// f = i.(float64) // creates panic
	// fmt.Println(f)

	TypeSwitchTest()
}

func do(i interface{}) {
	switch v := i.(type) {
	case string:
		{
			fmt.Printf("%v is string type\n", v)
		}
	case int:
		{
			fmt.Printf("%v is int type\n", v)
		}
	default:
		{
			fmt.Printf("I don't know about type %T\n", v)
		}
	}
}

func TypeSwitchTest() {
	do(nil)
	do(21)
	do("hello")
	do(true)
}
