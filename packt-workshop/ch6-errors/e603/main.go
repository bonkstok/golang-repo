package main

import (
	"errors"
	"fmt"
)

var (
	ErrHourlyRate  = errors.New("invalid hourly rate")
	ErrHoursWorked = errors.New("invalid hours worked per week")
)

func main() {

	pay, err := payDay(57, 50)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pay)

	pay, err = payDay(57, 100)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pay)

	pay, err = payDay(100, 50)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pay)

}

func payDay(hoursWorked, hourlyRate int) (int, error) {

	if hourlyRate < 10 || hourlyRate > 75 {
		return 0, ErrHourlyRate
	} else if hoursWorked < 0 || hoursWorked > 80 {
		return 0, ErrHoursWorked
	} else if hoursWorked > 40 {
		hoursOver := hoursWorked - 40
		overTime := hoursOver * 2
		regularPay := hoursWorked * hourlyRate
		return regularPay + overTime, nil
	} else {
		return hoursWorked * hourlyRate, nil
	}
}
