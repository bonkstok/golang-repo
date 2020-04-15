package main

import "fmt"

func main() {

	visitCount := 24
	visitCount++
	customerBill := [2]float64{13, 13}
	var total float64

	for i := 0; i < len(customerBill); i++ {
		total = total + customerBill[i]
	}
	fmt.Println("Sub:", total)

	total = total + (4 * 2.25)
	fmt.Println("Sub:", total)

	//discount yeaaah
	total = total - 5
	fmt.Println("Sub:", total)

	tip := total * 0.1
	fmt.Println("Tip:", tip)

	fmt.Println("Total:", (total + tip))

	remainder := visitCount % 5
	if remainder == 0 {
		fmt.Println("WHAT YOU WON A PRICE MAN!")
	}

}
