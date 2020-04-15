package main

import "fmt"

func message() string {
	arr := [...]string{
		"ready",
		"get",
		"go",
		"to",
	}
	return fmt.Sprintln(arr[1], arr[0], arr[3], arr[2])
}

func main() {
	a := message()
	//fmt.Print(message())
	fmt.Print(a)
}
