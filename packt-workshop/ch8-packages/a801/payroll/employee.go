package payroll

import (
	"fmt"
)

type Payer interface {
	Pay() (string, float64)
	GetFullName() string
}

type Employee struct {
	Id, FirstName, LastName string
}

func PayDetails(p Payer) {
	fmt.Println(p.Pay())
}

func GetFullName(p Payer) string {
	name := p.GetFullName()
	return name
}
