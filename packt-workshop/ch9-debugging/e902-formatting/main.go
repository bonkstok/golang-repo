package main

import "fmt"

func main() {
	for i := 1; i < 255; i++ {
		fmt.Printf("Decimal: %3.d Base Two: %8.b Hex: %2.x\n", i, i, i)
	}

	myArr := [5]int{}

	for k, _ := range myArr {
		fmt.Printf("\t%4.d\n", k)
	}
}
