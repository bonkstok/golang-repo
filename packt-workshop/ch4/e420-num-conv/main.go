package main

import (
	"fmt"
	"math"
)

func convert() string {
	var i8 int8 = math.MaxInt8
	i := 128
	f64 := 3.14

	//convert a smaller int into a alrger int - safe to do
	m := fmt.Sprintf("int8 = %v > int64 = %v\n", i8, int64(i8))

	//convert an int into a smaller int type - cause it to overflow
	m += fmt.Sprintf("int = %v > int8 = %v\n", i, int8(i))

	//convert int into float64
	m += fmt.Sprintf("int8 = %v > float64 = %v\n", i8, float64(i8))

	//covnvert float back to int
	m += fmt.Sprintf("Float64 = %v > int = %v\n", f64, int(f64))

	return m
}

func main() {
	fmt.Println(convert())
}
