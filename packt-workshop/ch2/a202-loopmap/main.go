package main

import "fmt"

func main() {
	words := map[string]int{
		"Gonna": 3,
		"You":   3,
		"Give":  2,
		"Never": 1,
		"Up":    4,
		"Bye":   10,
	}
	getMax(words)
}

func getMax(input map[string]int) {

	/*
		if len(input) == 0 {
			return errors.New("Error empty map")
		}
	*/
	var topWord string
	var topCount int = 0

	for key, value := range input {
		if value > topCount {
			topWord = key
			topCount = value
		}
		//fmt.Println(key, "--", value)
	}
	fmt.Printf("Top word %#v -- %#v", topWord, topCount)

}
