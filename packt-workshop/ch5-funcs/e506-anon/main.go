package main

import "fmt"

func main() {
	j := 9
	//this anon func does not have () at the end, this allows you to call it later like x(5) or x(j)
	x := func(i int) int {
		return i * i
	}

	fmt.Println(x(5))
	fmt.Println(x(j))

}
