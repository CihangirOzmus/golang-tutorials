package main

import "log"

type User struct{
	FirstName string
	LastName string
}

func main()  {
	for i := 0; i < 10; i++ {
		log.Println(i)
	}

	mySlice := []string{"dog", "cat", "horse", "fish", "banana"}
	for i, x := range mySlice {
		log.Println(i, x)
	}

	for _, x := range mySlice {
		log.Println(x)
	}

	myMap := make(map[string]string)
	myMap["Dog"] = "dog"
	myMap["Cat"] = "cat"
	myMap["Fish"] = "fish"
	myMap["hat"] = "hat"

	for i, v := range myMap {
		log.Println(i, v)
	}

	var myUserSlice []User
	u1 := User{
		FirstName: "Trevor",
		LastName: "Sawler",
	}

	u2 := User{
		FirstName: "Sam",
	}

	myUserSlice = append(myUserSlice, u1, u2)
	for i, v := range myUserSlice {
		log.Println(i, v)
	}
}