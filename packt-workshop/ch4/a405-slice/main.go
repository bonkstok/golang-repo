package main

import "fmt"

func main() {
	mySlice := []string{"Good", "Good", "Bad", "Good", "Good"}

	badIndex := len(mySlice) + 1 //atleast you will get an out of range

	for i := 0; i < len(mySlice); i++ {
		if mySlice[i] == "Bad" {
			badIndex = i
		}
	}

	mySlice = append(mySlice[:badIndex], mySlice[(badIndex+1):len(mySlice)]...)

	fmt.Println(mySlice)

}
