package main

import "fmt"

func main() {
	devSalary := salary(50, 2080, developerSalary)
	bossSalary := salary(150000, 25000, managerSalary)
	internSalary := salary(50, 2000, internSalary)
	fmt.Printf("Boss salary: %d\n", bossSalary)
	fmt.Printf("Developer salary: %d\n", devSalary)
	fmt.Printf("Intern salary: %d\n", internSalary)
}

func salary(x, y int, f func(int, int) int) int { // 2 parameters int, 1 param int that has signature 2 int params and returns int
	pay := f(x, y)
	return pay
}

func managerSalary(baseSalary, bonus int) int {
	return baseSalary + bonus
}

func developerSalary(hourlyRate, hoursWorked int) int {
	return hourlyRate * hoursWorked
}

func internSalary(_, _ int) int {
	return 0
}
