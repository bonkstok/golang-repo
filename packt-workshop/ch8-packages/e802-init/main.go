package main

import "fmt"

var budgetCategories = map[int]string{}

func init() {
	fmt.Println("Init our catagories...")
	budgetCategories[1] = "Car Insurance"
	budgetCategories[2] = "Mortgage"
	budgetCategories[3] = "Electricity"
	budgetCategories[4] = "Retirement"
	budgetCategories[5] = "Vacation"
	budgetCategories[7] = "Groceries"
	budgetCategories[8] = "Car Payment"
}

func main() {
	for k, v := range budgetCategories {
		fmt.Printf("Key: %v => Value: %s\n", k, v)
	}
}
