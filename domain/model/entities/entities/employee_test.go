package entities_test

import (
	"github.com/stretchr/testify/assert"
	"rest_app/domain/model/entities/entities"
	"testing"
)

func TestEmployee_CalculateAnnualSalary_HourlySalaryEmployee_OK(t *testing.T) {
	employeeMock := entities.Employee{
		HourlySalary:     10,
		ContractTypeName: entities.HourlySalaryEmployee,
	}
	expect := 120 * employeeMock.HourlySalary * 12
	employeeMock.CalculateAnnualSalary()
	assert.Equal(t, expect, employeeMock.AnnualSalary)
}

func TestEmployee_CalculateAnnualSalary_HourlySalaryEmployee_Error(t *testing.T) {
	employeeMock := entities.Employee{
		HourlySalary:     10,
		ContractTypeName: entities.HourlySalaryEmployee,
	}
	expect := 120 * employeeMock.HourlySalary * 10
	employeeMock.CalculateAnnualSalary()
	assert.NotEqual(t, expect, employeeMock.AnnualSalary)
}


func TestEmployee_CalculateAnnualSalary_MonthlySalaryEmployee_Ok(t *testing.T) {
	employeeMock := entities.Employee{
		MonthlySalary:     10,
		ContractTypeName: entities.MonthlySalaryEmployee,
	}
	expect := employeeMock.MonthlySalary * 12
	employeeMock.CalculateAnnualSalary()
	assert.Equal(t, expect, employeeMock.AnnualSalary)
}

func TestEmployee_CalculateAnnualSalary_MonthlySalaryEmployee_Error(t *testing.T) {
	employeeMock := entities.Employee{
		MonthlySalary:     10,
		ContractTypeName: entities.MonthlySalaryEmployee,
	}
	expect := employeeMock.MonthlySalary * 10
	employeeMock.CalculateAnnualSalary()
	assert.NotEqual(t, expect, employeeMock.AnnualSalary)
}
