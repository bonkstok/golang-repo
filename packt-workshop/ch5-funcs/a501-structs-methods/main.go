package main

import "fmt"

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type Employee struct {
	Id        string
	FirstName string
	LastName  string
}

type Developer struct {
	Individual Employee
	HourlyRate int
	WorkWeek   [7]int
}

func (d *Developer) LogHours(day Weekday, i int) {
	fmt.Println(day)
	d.WorkWeek[day] = i
	//fmt.Println(Weekday)
}

func (d *Developer) hoursWorked() int {
	totalHours := 0
	for _, v := range d.WorkWeek {
		totalHours += v
	}
	return totalHours
}

func (d *Developer) getHoursWorkedDay(day Weekday) int {
	return d.WorkWeek[day]
}

func (d *Developer) getHoursWorkedSummary() {
	fmt.Printf("Hours worked on Sunday %v\n", d.getHoursWorkedDay(Sunday))
	fmt.Printf("Hours worked on Monday %v\n", d.getHoursWorkedDay(Monday))
	fmt.Printf("Hours worked on Tuesday %v\n", d.getHoursWorkedDay(Tuesday))
	fmt.Printf("Hours worked on Wednesday %v\n", d.getHoursWorkedDay(Wednesday))
	fmt.Printf("Hours worked on Thursday %v\n", d.getHoursWorkedDay(Thursday))
	fmt.Printf("Hours worked on Friday %v\n", d.getHoursWorkedDay(Friday))
	fmt.Printf("Hours worked on Saturday %v\n", d.getHoursWorkedDay(Saturday))
}

func (d *Developer) totalEarnings() {
	totalIncome := 0
	for _, v := range d.WorkWeek {
		totalIncome += (v * d.HourlyRate)
	}
	fmt.Printf("Hours worked: %v\n", d.hoursWorked())
	fmt.Printf("Price/hour: %v\n", d.HourlyRate)
	fmt.Printf("Total income for %s,%s is: %v\n", d.Individual.LastName, d.Individual.FirstName, totalIncome)
}
func main() {
	dev1 := Developer{
		Individual: Employee{Id: "x007", FirstName: "Johnny", LastName: "Bond"},
		HourlyRate: 25,
	}
	dev1.LogHours(Monday, 5)
	dev1.LogHours(Friday, 15)
	fmt.Println(dev1.Individual.FirstName)
	fmt.Println(dev1.WorkWeek)
	fmt.Println(dev1.hoursWorked())
	dev1.totalEarnings()
	dev1.getHoursWorkedSummary()
}
