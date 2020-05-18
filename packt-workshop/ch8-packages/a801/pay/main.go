package main

import (
	"fmt"
	"payroll"
)

func main() {
	employeeReview := make(map[string]interface{})
	employeeReview["WorkQuality"] = 5
	employeeReview["TeamWork"] = 2
	employeeReview["Communication"] = "Poor"
	employeeReview["Problem-solving"] = 4
	employeeReview["Dependability"] = "Unsatisfactory"

	d1 := payroll.Developer{Individual: payroll.Employee{FirstName: "Johnny", LastName: "Test"}, HourlyRate: 15, HoursWorkedInYear: 1600, Review: employeeReview}
	m1 := payroll.Manager{Individual: payroll.Employee{FirstName: "Opper", LastName: "Manager"}, CommisionRate: 1.5, Salary: 65000}
	//err := d1.ReviewRating()
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	//fmt.Println(d1.GetFullName())
	fmt.Println(payroll.GetFullName(&d1))
	payroll.PayDetails(&d1)
	fmt.Println(payroll.GetFullName(&m1))
	payroll.PayDetails(&m1)
}
