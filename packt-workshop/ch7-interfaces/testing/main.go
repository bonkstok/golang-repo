package main

import (
	"fmt"
)

type Speaker interface {
	Speak() string
}
type cat struct {
	name string
	age  int
}

func main() {
	c := cat{name: "Oreo", age: 9}
	fmt.Println(c.Speakk())
	fmt.Println(c)
	myAnything("s")
	myAnything(4)
	myAnything([]int{1, 2, 4, 5})
	myRandom := []interface{}{1, "johnny", true}
	fmt.Println(myRandom)
	loopMe(myRandom)
}
func (c cat) Speakk() string {
	return "Purr Meow"
}
func (c cat) String() string {
	return fmt.Sprintf("%v (%v years old)", c.name, c.age)
}

func myAnything(s interface{}) {
	fmt.Println(s)
}

func loopMe(input []interface{}) {
	for _, v := range input {
		fmt.Printf("Element is of type %T and value %v\n", v, v)
	}
}
