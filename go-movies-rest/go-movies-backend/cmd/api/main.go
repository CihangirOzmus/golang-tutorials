package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

const version = "1.0.0"

type Config struct {
	port int
	env  string
}

func main() {
	var cfg Config

	flag.IntVar(&cfg.port, "port", 4000, "Server port to listen")
	flag.StringVar(&cfg.env, "env", "development", "Application environment")
	flag.Parse()

	fmt.Println("Running")

	http.HandleFunc("/status", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "status")
	})

	err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.port), nil)

	if err != nil {
		log.Println(err)
	}

}
