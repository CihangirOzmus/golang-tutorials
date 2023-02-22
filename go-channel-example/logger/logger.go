package logger

import (
	"log"
)

func ListenForLog(c chan string) {
	for {
		msg := <-c
		log.Println(msg)
	}
}
