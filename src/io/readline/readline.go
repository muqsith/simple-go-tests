package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	inFile := "/tmp/webshop-stats-workarea/merged/orders.json"
	outFile := "/tmp/out.json"
	// os
	fileReader, err := os.OpenFile(inFile, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer fileReader.Close()

	w, err := os.OpenFile(outFile, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer w.Close()

	lineSizeGuess := 1024 * 5000 // 5 MB

	bufferedReader := bufio.NewReaderSize(fileReader, lineSizeGuess)

	for line, _, err := bufferedReader.ReadLine(); err == nil && len(line) > 0; {
		w.Write(line)
		w.Write([]byte{'\n'})
		line, _, err = bufferedReader.ReadLine()
	}
}
