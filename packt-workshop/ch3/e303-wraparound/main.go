package main

import "fmt"

func main() {

	var a int8 = 125
	var b uint8 = 253

	for i := 0; i < 5; i++ {
		a++
		b++
		fmt.Printf("i: %v\n%T: %v\n%T: %v\n", i, a, a, b, b)
	}
}
