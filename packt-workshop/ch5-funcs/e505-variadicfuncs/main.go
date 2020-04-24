package main

import "fmt"

/*
func main() {
	myFunc(2)
	myFunc(2, 3)
	myFunc(2, 22, 41, 121)

	mySlice := []int{4, 5, 6, 7, 8, 9, 10}
	myFunc(mySlice...)

	myString("Hey", "alles", "goed")
}

func myFunc(i ...int) {
	fmt.Printf("Arguments given to this function: %v\t", len(i))
	fmt.Println(i)
}

func myString(s ...string) {
	fmt.Println(s)
}
*/

func main() {
	i := []int{5, 10, 15}
	fmt.Println(sum(5, 4))
	fmt.Println(sum(i...))
}

func sum(i ...int) int {
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}
