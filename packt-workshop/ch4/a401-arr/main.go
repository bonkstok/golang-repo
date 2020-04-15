package main

import "fmt"

func fillArray() [10]int {
	arr := [10]int{}

	for i := 0; i < len(arr); i++ {
		arr[i] = i + 1
	}

	return arr
}

func main() {
	fmt.Println(fillArray())
}
