package main

import (
	"bufio"
	"log"
	"os"
)

func CopyFiles() {
	inFile := "/home/muqsith/Development/simple-go-tests/tmp/files.zip"

	outFile := "/home/muqsith/Development/simple-go-tests/tmp/out.zip"
	os.Remove(outFile)

	fileReader, err := os.OpenFile(inFile, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer fileReader.Close()

	bufSize := 2 * 1000 * 1024

	bufferedReader := bufio.NewReaderSize(fileReader, bufSize)

	fileWriter, err := os.OpenFile(outFile, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer fileWriter.Close()

	buf := make([]byte, bufSize)

	for nread, err := bufferedReader.Read(buf); err == nil && nread > 0; {
		fileWriter.Write(buf[0:nread])
		nread, err = bufferedReader.Read(buf)
	}
}
