package main

import "fmt"

//Generally you define struct type in the package scope
type user struct {
	name    string
	age     int
	balance float64
	member  bool
}

func getUsers() []user {
	//order of defining doesnt matter, if you leave out any of the fields it will be assigned the zero value for that type
	u1 := user{
		name:    "Tracy",
		age:     51,
		balance: 2102.20,
		member:  false,
	}
	u2 := user{
		name: "Nick",
		age:  19,
	}
	// also possible to initialize a struct using only values -> the order is the order defined in the struct, all fieldnames must be specified
	// Bad readability, please dont use it
	u3 := user{
		"bob",
		25,
		0,
		false,
	}
	var u4 user //var notation will create a struct where all the fields have zero values

	u4.name, u4.age, u4.member, u4.balance = "Sue", 31, true, 17.09

	return []user{u1, u2, u3, u4}

}

func main() {
	users := getUsers()
	for i := 0; i < len(users); i++ {
		fmt.Printf("%v: %#v\n", i, users[i])
	}
}
