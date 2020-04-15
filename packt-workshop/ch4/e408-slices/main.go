package main

import (
	"fmt"
	"os"
)

func getPassedArgs(minArgs int) []string {
	if len(os.Args) < minArgs {
		fmt.Printf("Not enough arg(s). Need at least %v argument(s)", minArgs)
		os.Exit(1)
	}

	var args []string
	for i := 1; i < len(os.Args); i++ {
		args = append(args, os.Args[i])
	}

	return args
}

func findLongest(args []string) string {

	var longest string

	for _, index := range args {

		if len(index) > len(longest) {
			longest = index
		} else if len(index) == len(longest) {
			fmt.Println("Same size")
		} else {
			continue
		}
	}
	return longest
}

func main() {
	if longest := findLongest(getPassedArgs(3)); len(longest) > 0 {
		fmt.Println("The longest word passed was:", longest)
	} else {
		fmt.Println("There was an error")
		os.Exit(1)
	}

	fmt.Println(os.Args[0])
}
