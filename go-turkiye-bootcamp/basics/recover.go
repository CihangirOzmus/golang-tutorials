package basics

import (
	"fmt"
	"time"
)

func timer(c <-chan time.Time, message string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	for {
		select {
		case <-c:
			return
		default:
			fmt.Println(time.Now(), message)
			time.Sleep(1 * time.Second)

			if time.Now().Day() == 26 {
				panic("Hata oldu ya la")
			}
		}
	}
}
