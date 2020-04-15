package main

import "fmt"

var Y int = 10

func add5Value(count int) {
	count += 5
	fmt.Println("add5Value :", count)
}

func add5Point(count *int) {

	if count != nil {
		//count = &Y
		*count = *count + 5
		fmt.Println("add5Point :", *count)
	}
}

func main() {
	var count int
	var count2 *int
	add5Value(count)
	fmt.Println("add5Value post:", count) // this will print 0 since we passed a copy of the variable to the function

	count2 = &Y          //assign pointer to Y
	fmt.Println(*count2) //print value of Y

	add5Point(count2)                       // give address of Y as input
	fmt.Println("add5Point post:", *count2) // print current value of Y -> since in function we updated the pointer -> result is added to Y
	fmt.Println(Y)                          // same value as *count2 since its a pointer to Y
}
