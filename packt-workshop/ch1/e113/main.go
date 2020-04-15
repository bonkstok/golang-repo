package main

import (
	"fmt"
	"time"
)

func main() {
	var var1 *int // value will be nil
	//fmt.Print(&var1)
	fmt.Println(var1)
	var2 := new(int)
	fmt.Println(&var2)

	countTemp := 5
	fmt.Println(countTemp)
	t := time.Time{}
	fmt.Println(&t)

	*var2 = 10
	fmt.Println(var2) //this will have a different address than &var2 since you pass it to a function
	fmt.Println(&var2)
	fmt.Printf("Value of var2: %#v -- address of var2 %#p \n", *var2, &var2)

}
