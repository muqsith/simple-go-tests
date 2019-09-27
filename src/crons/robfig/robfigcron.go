package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/robfig/cron"
)

func runAtRegularIntervals() {
	fmt.Println("Running .... 	", time.Now())
}

func test() { fmt.Println("Every hour on the half hour") }

// courtesy : https://stackoverflow.com/a/28882045

func main() {
	c := cron.New()
	c.AddFunc("* * * * *", test)
	go c.Start()
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
}
