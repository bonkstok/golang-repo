package main

import (
	"fmt"
)

func main() {

	var count1 *int
	fmt.Println(&count1)
	count2 := new(int)
	countTemp := 5
	count3 := &countTemp
	//t := &time.Time{}
	fmt.Printf("Passing mem address %p\n", &count1)
	intCheck(count1)

	fmt.Printf("Passing mem address %p\n", &count2)
	intCheck(count2)

	fmt.Printf("Passing mem address %p\n", &count3)
	intCheck(count3)
}

func intCheck(x *int) {
	fmt.Printf("Received mem address %p \n", &x)
	fmt.Printf("Value of arg %#v\n", x)
	if x != nil {
		fmt.Println(*x)
	}
}
