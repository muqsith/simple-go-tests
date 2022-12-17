package main

import (
	"io"
	"os"
	"time"
)

func main() {
	io.WriteString(os.Stdout, "Hello,")
	time.Sleep(1 * time.Second)
	io.WriteString(os.Stdout, "\033[2K\r") // https://stackoverflow.com/a/1508589/2388706
	io.WriteString(os.Stdout, "World!")
	time.Sleep(1 * time.Second)
	io.WriteString(os.Stdout, "\033[2K\r")
	io.WriteString(os.Stdout, "\n")
}
