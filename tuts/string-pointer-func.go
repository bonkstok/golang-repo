/*
Checking how the pointers work when using them in functions
*/

package main

import (
	"fmt"
)


func main () {
	var myString string = "Johnny"

	fmt.Println(&myString) //print address
	strPoint(&myString)
	strFunc(myString)
	updateStrPoint(&myString)
	fmt.Println(&myString)
	fmt.Println(myString)
}

func strPoint(str *string){
	fmt.Println(str)
}

func updateStrPoint(str *string){
	*str = "updated"
}

func strFunc(str string){
	fmt.Println(&str)
}
