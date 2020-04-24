package main

import "fmt"

type Developer struct {
	Individual Employee
	HourlyRate int
	WorkWeek   [7]int
}
type Employee struct {
	Id        int
	FirstName string
	LastName  string
}
type Weekday int

const REGULARWORKWEEKHOURS int = 40

const (
	Sunday Weekday = iota //starts at zero
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	d := Developer{Individual: Employee{Id: 1, FirstName: "Tony", LastName: "Stark"}, HourlyRate: 10}
	x := nonLoggedHours()
	fmt.Println(x(4))
	d.LogHours(Monday, 8)
	d.LogHours(Tuesday, 8)
	d.LogHours(Wednesday, 8)
	d.LogHours(Friday, 16)
	d.LogHours(Saturday, 10)
	//fmt.Println(d.WorkWeek)
	d.PayDetails()

}

func nonLoggedHours() func(int) int {
	total := 0
	return func(i int) int {
		total += i
		return total
	}
}

func (d *Developer) LogHours(day Weekday, hours int) {
	d.WorkWeek[day] = hours
}

func (d *Developer) payDay() (int, bool) {
	totalHours := d.totalHoursWorkedWeek()
	if totalHours > REGULARWORKWEEKHOURS {
		regularPay := d.HourlyRate * REGULARWORKWEEKHOURS
		overTimePay := (d.HourlyRate * 2) * (totalHours - REGULARWORKWEEKHOURS)
		return regularPay + overTimePay, true
	}
	return totalHours * d.HourlyRate, false
}

func (d *Developer) totalHoursWorkedWeek() int {
	total := 0
	for _, v := range d.WorkWeek {
		total += v
	}
	return total
}

func (d *Developer) PayDetails() {
	for i, v := range d.WorkWeek {
		switch i {
		case 0:
			fmt.Println("Sunday hours: ", v)
		case 1:
			fmt.Println("Monday hours: ", v)
		case 2:
			fmt.Println("Tuesday hours: ", v)
		case 3:
			fmt.Println("Wednesday hours: ", v)
		case 4:
			fmt.Println("Thursday hours: ", v)
		case 5:
			fmt.Println("Friday hours: ", v)
		case 6:
			fmt.Println("Saturday hours: ", v)
		}
	}
	fmt.Printf("\nHours worked this week:  %d\n", d.totalHoursWorkedWeek())
	pay, overtime := d.payDay()
	fmt.Println("Pay for the week: $", pay)
	fmt.Println("Target hours:", REGULARWORKWEEKHOURS)
	fmt.Println("Is this overtime pay: ", overtime)
	if overtime {
		fmt.Println("Hours overtime:", (d.totalHoursWorkedWeek() - REGULARWORKWEEKHOURS))
	}
	fmt.Println()
}
