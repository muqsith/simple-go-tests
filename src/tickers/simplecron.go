package main

import (
	"fmt"
	"time"
)

func runAtRegularIntervals(t time.Time) {
	fmt.Println("Running .... 	", t)
}

func startCron() {
	ticker := time.NewTicker(500 * time.Millisecond)

	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				{
					runAtRegularIntervals(t)
				}
			}
		}
	}()

	time.Sleep(365 * 24 * 60 * 60 * time.Second)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}

func main() {
	startCron()
}
