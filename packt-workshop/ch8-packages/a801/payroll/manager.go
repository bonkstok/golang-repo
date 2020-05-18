package payroll

type Manager struct {
	Individual    Employee
	Salary        float64
	CommisionRate float64
}

func (m Manager) Pay() (string, float64) {
	//fullName := d.fullName()
	return GetFullName(m), m.Salary * m.CommisionRate
}

func (m Manager) GetFullName() string {
	return m.Individual.FirstName + " " + m.Individual.LastName
}
