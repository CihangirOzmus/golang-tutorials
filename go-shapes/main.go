package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type fileWriter struct {
}

func (f fileWriter) Write(bs []byte) (n int, err error) {
	fmt.Println(string(bs))
	return len(bs), nil
}

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatal("File name should be given.")
	}
	fileName := args[1]
	fmt.Println(fileName)
	open, err := os.Open(fileName)
	if err != nil {
		log.Fatal(fileName, "not found.")
	}

	fw := fileWriter{}
	io.Copy(fw, open)
}
