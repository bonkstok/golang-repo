package main

import (
	"errors"
	"fmt"
	"os"
)

type Employee struct {
	id, firstName, lastName string
}

type Developer struct {
	individual        Employee
	hoursWorkedInYear float64
	hourslyRate       float64
	review            map[string]interface{}
}

type Manager struct {
	individual     Employee
	Salary         float64
	commissionRate float64
}

type Payer interface {
	Pay() (string, float64)
}

func (d *Developer) Pay() (string, float64) {
	return d.getFullName(), (d.hoursWorkedInYear * d.hourslyRate)
}

func (d *Developer) getFullName() string {
	return d.individual.firstName + " " + d.individual.lastName
}

func (m *Manager) Pay() (string, float64) {
	return m.getFullName(), m.Salary + (m.Salary * m.commissionRate)
}

func (m *Manager) getFullName() string {
	return m.individual.firstName + " " + m.individual.lastName
}

func payDetails(p Payer) {
	fmt.Println(p.Pay())
}

func (d *Developer) ReviewRating() error {
	total := 0
	for _, v := range d.review {
		rating, err := OverallReview(v)
		if err != nil {
			return err
		}
		total += rating
	}
	averageRating := float64(total) / float64(len(d.review))
	fmt.Printf("%s got a review rating of %.2f\n", d.getFullName(), averageRating)
	return nil
}

func OverallReview(i interface{}) (int, error) {
	switch v := i.(type) {
	case int:
		return v, nil
	case string:
		rating, err := convertReviewToInt(v)
		if err != nil {
			return 0, err
		}
		return rating, nil
	default:
		return 0, errors.New("Unknown type")
	}
}

func convertReviewToInt(str string) (int, error) {
	switch str {
	case "Excellent":
		return 5, nil
	case "Good":
		return 4, nil
	case "Fair":
		return 3, nil
	case "Poor":
		return 2, nil
	case "Unsatisfactory":
		return 1, nil
	default:
		return 0, errors.New("invalid rating: " + str)
	}
}

func main() {
	employeeReview := make(map[string]interface{})
	employeeReview["WorkQuality"] = 5
	employeeReview["TeamWork"] = 2
	employeeReview["Communication"] = "Poor"
	employeeReview["Problem-solving"] = 4
	employeeReview["Dependability"] = "Unsatisfactory"

	d1 := Developer{individual: Employee{firstName: "Johnny", lastName: "Test"}, hourslyRate: 15, hoursWorkedInYear: 1600, review: employeeReview}
	err := d1.ReviewRating()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//fmt.Println(d1.Pay())
	payDetails(&d1)
	m1 := Manager{individual: Employee{firstName: "Hello", lastName: "Manager"}, commissionRate: 1, Salary: 65000}
	payDetails(&m1)
}
