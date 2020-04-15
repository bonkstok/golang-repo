package main

import "fmt"

func fillArray(arr [10]int) [10]int {

	for i := 0; i < len(arr); i++ {
		arr[i] = i + 1
	}
	return arr
}

func opArray(arr [10]int) [10]int {

	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * arr[i]
	}
	return arr
}

func main() {
	arr1 := [10]int{1, 44, 5, 666, 7, 4, 31, 22, 11, 999}
	fmt.Println(fillArray(arr1))
	fmt.Println(opArray(arr1))
}
