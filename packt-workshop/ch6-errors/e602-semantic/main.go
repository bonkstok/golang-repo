package main

import "fmt"

const WALKLIMIT int = 2

func main() {
	km := 3

	//point is to take the car if km equal or greater than 2 (WALKLIMIT) using km > means it will only display if km is 3+
	if km > WALKLIMIT {
		fmt.Println("Take the car")
	} else {
		fmt.Println("Lets walk")
	}
}
