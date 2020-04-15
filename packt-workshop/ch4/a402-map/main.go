package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func getArgs() ([]int, error) {
	var args []int

	if len(os.Args) < 2 {
		//fmt.Println("Please give one or more keys to lookup.")
		return args, errors.New("Did not receive key to lookup")
	}

	for i := 1; i < len(os.Args); i++ {
		arg, err := strconv.Atoi(os.Args[i]) //try to convert string (argument) to integer
		if err != nil {
			fmt.Printf("Argument '%v' is not convertable to integer..Skipping value\n", os.Args[i])
			continue
		}
		args = append(args, arg)
	}

	return args, nil
}

func printKey(input map[int]string) {
	//first check the args passed
	args, err := getArgs()

	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(args); i++ {
		for k, v := range input {
			if args[i] == k {
				fmt.Println("Hello there,", v)
			}
		}
	}

}

func main() {
	myMap := map[int]string{
		305: "Sue",
		204: "Bob",
		631: "Jake",
		073: "Tracy",
		999: "Tina",
		555: "Dolly",
	}
	printKey(myMap)
}
