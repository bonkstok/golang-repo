package main

import "fmt"

func main() {
	/*
		counter := 4
		x := decrement(counter)
		fmt.Println(x())
		fmt.Println(x())
		fmt.Println(x())
		fmt.Println(x())
	*/
	number := 20
	//y := checker(number)           //check value
	z := decrement(number)         //decrement by one
	for i := number; i != 0; i-- { // set i to value of y doesnt print 20 because decrement func decrements before returning the value
		fmt.Println(z())
	}
}

func decrement(counter int) func() int {
	//counter := 5
	return func() int {
		counter--
		return counter
	}
}

func checker(counter int) func() int { //not needed.. pff me sometimes
	return func() int {
		return counter
	}
}
