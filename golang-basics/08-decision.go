package main

import "log"

func main()  {
	var isTrue bool
	isTrue = true
	if isTrue {
		log.Println("isTrue is", isTrue)
	} else {
		log.Println("isTrue is", isTrue)
	}

	cat := "cat"
	if cat == "cat" {
		log.Println("cat is cat")
	} else {
		log.Println("cat is not cat")
	}

	myNum := 90
	isTrue = false

	if myNum > 99 && isTrue {
		log.Println("myNum is greater than 99 and isTrue set to true")
	} else if myNum < 100 && isTrue {
		log.Println("myNum is less than 99 and isTrue set to true")
	} else if myNum == 100 || isTrue {
		log.Println("myNum is equal to 100 or isTrue set to true")
	} else if myNum > 1000 && !isTrue {
		log.Println("myNum is greater than 1000 and isTrue set to false")
	}

}