package basics

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type UserInterface interface {
	getFollowerSize() int
	addFollower(u user) error
}

type user struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       uint8  `json:"age"`
	Followers []user `json:"followers,omitempty"`
}

func main() {
	user1 := user{
		FirstName: "Cihangir",
		LastName:  "Ozmus",
		Age:       40,
		Followers: []user{
			{
				FirstName: "Ahmet",
				LastName:  "Ozmus",
				Age:       23,
			},
		},
	}

	follower := user{
		FirstName: "Nisa",
		LastName:  "Ozmus",
		Age:       25,
	}

	println(user1.getFollowerSize())

	if err := user1.addFollower(follower); err != nil {
		log.Fatal(err)
	}

	println(user1.getFollowerSize())

	marshal, _ := json.Marshal(user1)
	fmt.Println(string(marshal))

	println(compareFollowerSize(&user1, &follower))
	println(compareFollowerSize(&follower, &user1))
}

func (u *user) getFollowerSize() int {
	return len(u.Followers)
}

func (u *user) addFollower(f user) error {
	if u.getFollowerSize() == 10 {
		return errors.New("max follower size")
	}

	u.Followers = append(u.Followers, f)
	return nil
}

func compareFollowerSize(u1, u2 UserInterface) bool {
	return u1.getFollowerSize() > u2.getFollowerSize()
}
