package main

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidRoutingNumber = errors.New("invalid routing number")
	ErrInvalidLastName      = errors.New("invalid last name")
)

type directDeposit struct {
	lastName, firstName, bankName string
	routingNumber, accountNumber  int
}

func (d *directDeposit) validateRoutingNumber() {
	defer d.report()
	if d.routingNumber < 100 {
		panic(ErrInvalidRoutingNumber)
	}
}

func (d *directDeposit) validateLastName() {
	defer d.report()
	if d.lastName == "" {
		panic(ErrInvalidLastName)
	}
}

func (d *directDeposit) report() {
	fmt.Println("Last name:", d.lastName)
	fmt.Println("First name:", d.firstName)
	fmt.Println("Bank name:", d.bankName)
	fmt.Println("Routing number:", d.routingNumber)
	fmt.Println("Account number:", d.accountNumber)
}
func main() {
	customer1 := directDeposit{firstName: "Johnny", lastName: "", bankName: "Underground", routingNumber: 919, accountNumber: 201232}

	customer1.validateRoutingNumber()
	customer1.validateLastName()

	customer1.report()
}
