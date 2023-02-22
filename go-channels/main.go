package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://www.google.com/",
		"http://www.facebook.com/",
		"http://www.stackoverflow.com/",
		"http://www.golang.com/",
		"http://www.spotify.com/",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func(l string) {
			time.Sleep(5 * time.Second)
			checkLink(l, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	bt := time.Now()
	resp, err := http.Get(link)
	if err != nil {
		c <- link
	}
	elapsedTime := time.Now().Sub(bt).Milliseconds()
	fmt.Println(resp.Status, elapsedTime, "ms", link, " ")
	c <- link
}
