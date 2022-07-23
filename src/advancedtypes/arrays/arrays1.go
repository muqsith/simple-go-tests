package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	m1 := make(map[string]interface{})
	m1["list"] = []bool{}

	for i := 0; i < 10; i++ {
		list := m1["list"].([]bool)
		randomNumber := rand.Intn(500)
		boolValue := true
		if randomNumber%2 == 0 {
			boolValue = false
		}
		list = append(list, boolValue)
		m1["list"] = list
	}

	for i, v := range m1["list"].([]bool) {
		fmt.Println("Value at ", i, " is ", v)
	}

}
