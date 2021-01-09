package main

import (
	"encoding/json"
	"fmt"
)

type aPerson struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Age       int    `json:"Age"`
}

func main() {
	s := `[{"FirstName":"James","LastName":"Bond","Age":42},{"FirstName":"Jane","LastName":"Doe","Age":32}]`
	bs := []byte(s)
	fmt.Println(bs)

	var people []aPerson
	err := json.Unmarshal(bs, &people)

	if err != nil {
		fmt.Println(err)
	}

	for i, v := range people {
		fmt.Println(i, v)
	}
}
