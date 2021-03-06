package main

import "log"

type User struct{
	FirstName string
	LastName string
}

//maps are not sorted, do not keep order. No need to use pointers for map. Map always reference itself. Map is immutable
//also not recommend but map[string]interface{} can be used to store anything wanted
func main()  {
	//var myOther map[string]string

	myMap := make(map[string]string)//declare this way
	myMap["dog"] = "Samson"
	myMap["otherDog"] = "Cassie"

	log.Println(myMap["dog"], myMap["otherDog"])

	myMapInt := make(map[string]int)
	myMapInt["First"] = 1
	myMapInt["Second"] = 2

	log.Println(myMapInt["First"], myMapInt["Second"])

	myMapUser := make(map[string]User)
	me := User{
		FirstName: "Trevor",
		LastName: "Sawler",
	}
	myMapUser["me"] = me

	log.Println(myMapUser["me"].FirstName)
}