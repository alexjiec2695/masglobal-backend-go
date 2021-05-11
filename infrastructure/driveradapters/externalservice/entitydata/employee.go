package entitydata

type EmployeeData struct {
	Id               int     `json:"id"`
	Name             string  `json:"name"`
	ContractTypeName string  `json:"contractTypeName"`
	RoleId           int     `json:"roleId"`
	RoleName         string  `json:"roleName"`
	RoleDescription  string  `json:"roleDescription"`
	HourlySalary     float64 `json:"hourlySalary"`
	MonthlySalary    float64 `json:"monthlySalary"`
}
