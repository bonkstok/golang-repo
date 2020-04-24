package main

import "fmt"

var (
	budgetCategories = map[int]string{}
	payeeToCategory  = map[string]int{}
)

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

func init() {
	fmt.Println("Assign our Payees to categories")
	payeeToCategory["Nationwide"] = 1
	payeeToCategory["BBT Loan"] = 2
	payeeToCategory["First Energy Electric"] = 3
	payeeToCategory["Ameriprise Financial"] = 4
	payeeToCategory["Walt Disney World"] = 5
	payeeToCategory["ALDI"] = 7
	payeeToCategory["Martins"] = 7
	payeeToCategory["Wal Mart"] = 7
	payeeToCategory["Chevy Loan"] = 8
}
func main() {
	for k, v := range payeeToCategory {
		fmt.Printf("Payee: %s, category: %v\n", k, budgetCategories[v])
	}
}
