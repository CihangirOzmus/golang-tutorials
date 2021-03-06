package main

import "log"

type myStruct struct{
	FirstName string
}

// receiver <funcName> return type
func (m *myStruct) printFirstName() string {
	return m.FirstName
}

func main()  {
	var myVar myStruct
	myVar.FirstName = "John"

	myVar2 := myStruct{
		FirstName: "Marry",
	}

	log.Println("myVar is", myVar.FirstName)
	log.Println("myVar2 is", myVar2.FirstName)

	log.Println("myVar is", myVar.printFirstName())
	log.Println("myVar2 is", myVar2.printFirstName())
}