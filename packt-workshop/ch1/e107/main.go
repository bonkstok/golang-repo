package main

import (
	"fmt"
)

func printString(msg int) {
	fmt.Println(msg)
}

func main() {

	offset := 5
	printString(offset)
	offset = 10
	printString(offset)

}
