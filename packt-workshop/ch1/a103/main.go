package main

import "fmt"

func main() {
	count := 6

	var message string

	if count > 5 {
		message = "Count greater than 5"
	} else {
		message = "Count not greather than 5"
	}
	fmt.Println(message)
}
