package main

import "fmt"

type Person struct {
	name string
	age  int
	city string
}

func (person *Person) Print() {
	fmt.Println("Name: ", person.name, ", age: ", person.age, ", city: ", person.city)
}

func (person *Person) Age(years int) {
	person.age += years
}

func (person Person) String() string {
	return fmt.Sprintf("{\"name\": \"%v\", \"age\": %v, \"city\": \"%v\"}", person.name, person.age, person.city)
}

func main() {
	person1 := Person{
		name: "Irfan",
		age:  37,
		city: "Dubai",
	}

	person1.Age(10)

	fmt.Println(person1)
}
