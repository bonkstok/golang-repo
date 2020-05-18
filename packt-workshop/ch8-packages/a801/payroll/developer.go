package payroll

type Developer struct {
	Individual        Employee
	HourlyRate        float64
	HoursWorkedInYear float64
	Review            map[string]interface{}
}

func (d Developer) Pay() (string, float64) {
	//fullName := d.fullName()
	return GetFullName(&d), d.HourlyRate * d.HoursWorkedInYear
}

func (d *Developer) GetFullName() string {
	return d.Individual.FirstName + " " + d.Individual.LastName
}
