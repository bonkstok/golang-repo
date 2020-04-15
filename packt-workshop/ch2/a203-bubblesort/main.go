package main

import "fmt"

func main() {
	mySlice := []int{8, 2, 5, 6, 3, 4, 10, 1, 9, 7}
	//myOrderedSlice := []int{}

	for swapped := true; swapped; {
		swapped = false
		for i := 1; i < len(mySlice); i++ {
			fmt.Println("is:", i, "-- current:", mySlice[i], "myslice[-1]", mySlice[i-1], "First element is", mySlice[0])
			if mySlice[i] < mySlice[i-1] {

				mySlice[i], mySlice[i-1] = mySlice[i-1], mySlice[i]
				swapped = true
			}
		}
	}
	fmt.Println(mySlice)
	fmt.Println(mySlice[len(mySlice)-1])
	fmt.Println(mySlice[6-1])
	mySlice[0] = 5
	fmt.Println(mySlice)
}
