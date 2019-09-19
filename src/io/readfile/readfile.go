package main

import (
	"fmt"
	"os"
)

func main() {
	const SRCPATH = "/tmp/webshop-stats-workarea/merged/order_events.json"
	const DSTPATH = "/tmp/outfile"

	fmt.Printf("Source file: \n%s\n", SRCPATH)
	src, err := os.OpenFile(SRCPATH, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer src.Close()

	dst, err := os.OpenFile(DSTPATH, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Destination file: \n%s\n", DSTPATH)
	defer dst.Close()

	buf := make([]byte, 1024*10000)

	for nread, err := src.Read(buf); err == nil && nread > 0; {
		dst.Write(buf[0:nread])
		nread, err = src.Read(buf)
	}
}
