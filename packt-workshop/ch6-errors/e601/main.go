package main

import "fmt"

func main() {
	testArr := [5]int{1, 2, 3, 4, 5}

	//this will cause any error since testArr is only of size 5 - panic: runtime error: index out of range
	//for i := 0; i < 10; i++ {
	//	fmt.Println(testArr[i])
	//}

	//can be resolved by using the range
	for _, v := range testArr {
		fmt.Println(v)
	}
}
