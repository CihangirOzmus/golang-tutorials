package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type logWriter struct {
}

func (l logWriter) Write(bs []byte) (n int, err error) {
	fmt.Println(string(bs))
	return len(bs), nil
}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		log.Fatal("Could not get the response!")
	}

	lw := logWriter{}

	io.Copy(lw, resp.Body)
}
