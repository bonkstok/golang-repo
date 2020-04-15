package main

import "fmt"

const (
	x = string(iota + 'a')
	y
	z
)

func main() {
	count := 0
	if count < 5 {
		count = 10
		count++
	}
	fmt.Println(count == 11)
	fmt.Println(z)
}
