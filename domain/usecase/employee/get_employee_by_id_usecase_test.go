package employee_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"rest_app/domain/model/entities/entities"
	"rest_app/domain/usecase/employee"
	"testing"
)

func TestGetEmployeeByIdUseCase_GetEmployeeById_OK(t *testing.T) {
	employeeGateway := new(EmployeeGatewayMock)

	employeeMock := entities.Employee{
		HourlySalary:     10,
		MonthlySalary:    10,
		ContractTypeName: entities.MonthlySalaryEmployee,
	}

	employeeExpect := entities.Employee{
		HourlySalary:     10,
		MonthlySalary:    10,
		ContractTypeName: entities.MonthlySalaryEmployee,
		AnnualSalary:     120,
	}

	employeeGateway.On("GetEmployeeById", 1).Return(employeeMock, nil)
	newUseCase := employee.NewGetEmployeeByIdUseCase(employeeGateway)
	employeeResult, err := newUseCase.GetEmployeeById(1)

	assert.Equal(t, employeeExpect, employeeResult)
	assert.NoError(t, err)
}

func TestGetEmployeeByIdUseCase_GetEmployeeById_Error(t *testing.T) {
	employeeGateway := new(EmployeeGatewayMock)

	employeeGateway.On("GetEmployeeById", 1).Return(entities.Employee{}, errors.New("error"))
	newUseCase := employee.NewGetEmployeeByIdUseCase(employeeGateway)
	_, err := newUseCase.GetEmployeeById(1)

	assert.Error(t, err)
}
