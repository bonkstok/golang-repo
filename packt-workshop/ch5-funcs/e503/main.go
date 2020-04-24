package main

import "fmt"

func main() {
	for i := 1; i <= 15; i++ {
		n, s := fizzBuzz(i)
		fmt.Printf("Results: %d %s\n", n, s)
	}
}

func fizzBuzz(input int) (int, string) {
	switch {
	case input%15 == 0:
		return input, "FizzBuzz"
	case input%3 == 0:
		return input, "Fizz"
	case input%5 == 0:
		return input, "Buzz"
	default:
		return input, ""
	}
}
