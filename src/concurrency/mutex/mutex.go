package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

func (counter *SafeCounter) Inc(key string) {
	counter.mux.Lock()
	counter.v[key] += 1
	counter.mux.Unlock()
}

func (counter *SafeCounter) Value(key string) int {
	counter.mux.Lock()
	defer counter.mux.Unlock()
	return counter.v[key]
}

func main() {
	counter := SafeCounter{
		v: make(map[string]int),
	}
	for i := 0; i < 1000; i += 1 {
		go counter.Inc("somekey")
	}

	time.Sleep(time.Second * 3)
	fmt.Println(counter.Value("somekey"))
}
