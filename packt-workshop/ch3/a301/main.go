package main

import (
	"errors"
	"fmt"
)

func calculateTax(productPrice float32, taxRate float32) (float32, error) {

	if productPrice > 0 {
		return productPrice * taxRate, nil
	}

	return 0, errors.New("Could not calculate tax")
}

func main() {
	/*	shopCart := map[float32]float32{
			1.84:   .21,
			2.89:   .21,
			102.29: .21,
			0:      .21,
		}
	*/
	shopCart := map[string]map[float32]float32{
		"Bad milk":       {1.84: .21},
		"Smelly cheese":  {4.32: .21},
		"Weird stickers": {1.2: 0.21},
		"Socks":          {3.85: 0.21},
	}

	var totalTax, totalPrice float32 = 0, 0

	for key, value := range shopCart {
		for k, v := range value {
			tax, err := calculateTax(k, v)
			if err != nil {
				fmt.Println("Error calculating tax:", err)
				continue
			}
			fmt.Printf("Calculating price for product %v - price %v - taxrate %v - total price = %v\n", key, k, v, (k + tax))
			totalTax = totalTax + tax
			totalPrice = totalPrice + k
		}
	}
	/*
		for k, v := range shopCart {

			tax, err := calculateTax(k, v)
			if err != nil {
				fmt.Println("Error calculating tax:", err)
				continue
			}
			totalTax = totalTax + tax
			totalPrice = totalPrice + k
		}
	*/
	fmt.Println("Total untax:", totalPrice)
	fmt.Println("Total tax:", totalTax)
	fmt.Printf("Total: %.2f\n", totalPrice+totalTax)
}
