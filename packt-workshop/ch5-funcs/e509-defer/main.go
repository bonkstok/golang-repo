package main

import "fmt"

func main() {
	defer fmt.Println("hey from defer")
	fmt.Println("m1")
	saySomething()
	deferFunc()
}

func saySomething() {
	fmt.Println("Something")
}

func deferFunc() {
	defer fmt.Println("Im in a function ~ Defer")
	fmt.Println("Just printing to stdout from a function..")
}
