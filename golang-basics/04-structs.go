package main

import (
	"log"
	"time"
)

type User struct {
	FirstName string //capital letter makes is public
	LastName string
	PhoneNumber string
	Age int
	Birthday time.Time
}

func main()  {
	user := User {
		FirstName: "Trevor",
		LastName: "Sawler",
		PhoneNumber: "1 555 555 1212",
		Birthday: time.Date(1980, 1, 1, 0,0,0,0,time.UTC),
	}

	log.Println(user.FirstName, user.LastName, "BirthDate:", user.Birthday)
}