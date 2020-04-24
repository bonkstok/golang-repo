package main

import (
	"errors"
	"fmt"
)

var (
	ErrHourlyRate  = errors.New("invalid hourly rate")
	ErrHoursWorked = errors.New("invalid hours worked per week")
)

type Employee struct {
	firstName, lastName string
	hourlyRate          int
	hoursWorked         int
}

func main() {

	e1 := Employee{firstName: "Johnny", lastName: "van Veen", hourlyRate: 15, hoursWorked: 38}
	fmt.Printf("Summary for employee %s %s\nHours worked: %d\nHourly rate:%d\nWeekpay:%d\n", e1.firstName, e1.lastName, e1.hoursWorked, e1.hourlyRate, payDay(e1.hoursWorked, e1.hourlyRate))

	pay := payDay(40, 50)
	fmt.Println(pay)
	pay = payDay(81, 50)
	fmt.Println(pay)
}

func payDay(hoursWorked, hourlyRate int) int {
	report := func() {
		fmt.Printf("HoursWorked: %d\nHourslyRate: %d\n", hoursWorked, hourlyRate)
	}
	defer report()

	if hourlyRate < 10 || hourlyRate > 75 {
		panic(ErrHourlyRate)
	} else if hoursWorked < 0 || hoursWorked > 80 {
		panic(ErrHoursWorked)
	} else if hoursWorked > 40 {
		hoursOver := hoursWorked - 40
		overTime := hoursOver * 2
		regularPay := hoursWorked * hourlyRate
		return regularPay + overTime
	} else {
		return hoursWorked * hourlyRate
	}
}
