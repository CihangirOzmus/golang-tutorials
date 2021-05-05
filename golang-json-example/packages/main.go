package main

import (
	"encoding/json"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	myJson := `[{
		"first_name": "Clark",
		"last_name": "Kent",
		"hair_color": "black",
		"has_dog": true
	}, {
		"first_name": "Peter",
		"last_name": "Parker",
		"hair_color": "brown",
		"has_dog": false
	}]`

	var personList []Person

	err := json.Unmarshal([]byte(myJson), &personList)

	if err != nil {
		log.Println("Error at unmarshalling", err)
	}

	log.Printf("personList: %v", personList)

	var m1 Person = Person{FirstName: "Wally", LastName: "West", HairColor: "Red", HasDog: false}

	personList = append(personList, m1)

	newJson, err := json.MarshalIndent(personList, "", "    ")

	if err != nil {
		log.Println("Error marshalling", err)
	}

	log.Println(string(newJson))
}
