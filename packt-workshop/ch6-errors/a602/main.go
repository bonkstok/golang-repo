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

func (d *directDeposit) validateRoutingNumber() error {
	if d.routingNumber < 100 {
		return ErrInvalidRoutingNumber
	}
	return nil
}

func (d *directDeposit) validateLastName() error {
	if d.lastName == "" {
		return ErrInvalidLastName
	}
	return nil
}

func (d *directDeposit) report() {
	fmt.Println("Last name:", d.lastName)
	fmt.Println("First name:", d.firstName)
	fmt.Println("Bank name:", d.bankName)
	fmt.Println("Routing number:", d.routingNumber)
	fmt.Println("Account number:", d.accountNumber)
}
func main() {
	customer1 := directDeposit{firstName: "Johnny", lastName: "test", bankName: "Underground", routingNumber: 99, accountNumber: 201232}

	err := customer1.validateRoutingNumber()
	if err != nil {
		fmt.Println(err)
	}
	err = customer1.validateLastName()
	if err != nil {
		fmt.Println(err)
	}

	customer1.report()
}
