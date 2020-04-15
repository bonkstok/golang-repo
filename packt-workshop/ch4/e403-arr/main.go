package main

import "fmt"

func compArrays() (bool, bool, [10]int) {
	var arr1 [10]int
	arr2 := [...]int{9: 0}                 //
	arr3 := [10]int{1, 9: 10, 4: 5, 3: 52} // <index>: <value> so 9: 10 means that the 10 ele will have value 10 - start from 0

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)

	return arr1 == arr2, arr1 == arr3, arr3
}

func main() {
	comp1, comp2, arr3 := compArrays()
	fmt.Println("[10]int == [...]{9:0}       :", comp1)
	fmt.Println("[10]int == [10]int{1, 9: 10, 4: 5}}:", comp2)
	fmt.Println(arr3)
}
