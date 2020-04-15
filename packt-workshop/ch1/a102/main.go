package main

import "fmt"

func main() {
	a, b := 5, 10
	fmt.Println("before a", a)
	fmt.Println("before b", b)
	intSwap(&a, &b)
	fmt.Println("after a", a)
	fmt.Println("after b", b)

}

func intSwap(a *int, b *int) {
	var temp int = 0
	temp = *a
	*a = *b
	*b = temp
}
