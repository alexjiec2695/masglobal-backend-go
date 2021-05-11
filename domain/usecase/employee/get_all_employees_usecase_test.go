package employee_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"rest_app/domain/model/entities/entities"
	"rest_app/domain/usecase/employee"
	"testing"
)

type EmployeeGatewayMock struct {
	mock.Mock
}

func (mock *EmployeeGatewayMock) GetAllEmployees() ([]entities.Employee, error) {
	args := mock.Called()
	return args.Get(0).([]entities.Employee), args.Error(1)
}

func (mock *EmployeeGatewayMock) GetEmployeeById(ID int) (entities.Employee, error) {
	args := mock.Called(ID)
	return args.Get(0).(entities.Employee), args.Error(1)
}

func TestGetAllEmployeesUseCase_GetAllEmployees_OK(t *testing.T) {
	employeeGateway := new(EmployeeGatewayMock)

	employeeMock := make([]entities.Employee, 1)
	employeeMock[0] = entities.Employee{
		HourlySalary:     10,
		MonthlySalary:    10,
		ContractTypeName: entities.MonthlySalaryEmployee,
	}

	employeeExpect := make([]entities.Employee, 1)
	employeeExpect[0] = entities.Employee{
		HourlySalary:     10,
		MonthlySalary:    10,
		ContractTypeName: entities.MonthlySalaryEmployee,
		AnnualSalary:     120,
	}

	employeeGateway.On("GetAllEmployees").Return(employeeMock, nil)
	newUseCase := employee.NewGetAllEmployeesUseCase(employeeGateway)
	employeeResult, err := newUseCase.GetAllEmployees()

	assert.Equal(t, employeeExpect, employeeResult)
	assert.NoError(t, err)
}

func TestGetAllEmployeesUseCase_GetAllEmployees_Error(t *testing.T) {
	employeeGateway := new(EmployeeGatewayMock)

	employeeGateway.On("GetAllEmployees").Return([]entities.Employee{}, errors.New("error"))
	newUseCase := employee.NewGetAllEmployeesUseCase(employeeGateway)
	_, err := newUseCase.GetAllEmployees()

	assert.Error(t, err)
}


