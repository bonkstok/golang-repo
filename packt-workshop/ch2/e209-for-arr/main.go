package main

import "fmt"

func main() {
	names := []string{"Johnny", "Tina", "Roos", "Casper", "dolly"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
	for _, value := range names {
		//fmt.Println(index)
		fmt.Println(value)
	}

}
