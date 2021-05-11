package entities

type Employee struct {
	Id               int
	Name             string
	ContractTypeName string
	RoleId           int
	RoleName         string
	RoleDescription  string
	HourlySalary     float64
	MonthlySalary    float64
	AnnualSalary     float64
}

const (
	HourlySalaryEmployee  = "HourlySalaryEmployee"
	MonthlySalaryEmployee = "MonthlySalaryEmployee"
)

func (e *Employee) CalculateAnnualSalary() {

	if e.ContractTypeName == HourlySalaryEmployee {
		e.AnnualSalary = 120 * e.HourlySalary * 12
	}

	if e.ContractTypeName == MonthlySalaryEmployee {
		e.AnnualSalary = e.MonthlySalary * 12
	}
}
