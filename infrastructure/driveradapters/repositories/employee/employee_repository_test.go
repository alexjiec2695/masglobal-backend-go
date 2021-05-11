package employee_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"rest_app/domain/model/entities/entities"
	"rest_app/infrastructure/drivenadapters/repositories/employee"
	"testing"
)

type EmployeeAdapterMock struct {
	mock.Mock
}

func (mock *EmployeeAdapterMock) GetAllEmployees() ([]entities.Employee, error) {
	args := mock.Called()
	return args.Get(0).([]entities.Employee), args.Error(1)
}

func (mock *EmployeeAdapterMock) GetEmployeeById(ID int) (entities.Employee, error) {
	args := mock.Called(ID)
	return args.Get(0).(entities.Employee), args.Error(1)
}

func TestRepository_GetAllEmployees_OK(t *testing.T) {
	employeeAdapterMock := new(EmployeeAdapterMock)

	employeeAdapterMock.On("GetAllEmployees").Return([]entities.Employee{}, nil)

	newRepository := employee.NewEmployeeRepository(employeeAdapterMock)

	employeeResult, err := newRepository.GetAllEmployees()

	assert.Equal(t, []entities.Employee{}, employeeResult)
	assert.NoError(t, err)
}

func TestRepository_GetAllEmployees_Error(t *testing.T) {
	employeeAdapterMock := new(EmployeeAdapterMock)

	employeeAdapterMock.On("GetAllEmployees").Return([]entities.Employee{}, errors.New("error"))

	newRepository := employee.NewEmployeeRepository(employeeAdapterMock)

	_, err := newRepository.GetAllEmployees()

	assert.Error(t, err)
}

func TestRepository_GetEmployeeById_OK(t *testing.T) {
	employeeAdapterMock := new(EmployeeAdapterMock)

	employeeReturn := make([]entities.Employee, 3)

	for i := 0; i < 3; i++ {
		employeeReturn[i] = entities.Employee{
			Id: i,
		}
	}

	employeeAdapterMock.On("GetAllEmployees").Return(employeeReturn, nil)

	newRepository := employee.NewEmployeeRepository(employeeAdapterMock)

	employeeResult, err := newRepository.GetEmployeeById(1)

	assert.Equal(t, entities.Employee{Id: 1}, employeeResult)
	assert.NoError(t, err)
}

func TestRepository_GetEmployeeById_Empty_OK(t *testing.T) {
	employeeAdapterMock := new(EmployeeAdapterMock)

	employeeAdapterMock.On("GetAllEmployees").Return([]entities.Employee{}, nil)

	newRepository := employee.NewEmployeeRepository(employeeAdapterMock)

	employeeResult, err := newRepository.GetEmployeeById(1)

	assert.Equal(t, entities.Employee{}, employeeResult)
	assert.NoError(t, err)
}

func TestRepository_GetEmployeeById_Error(t *testing.T) {
	employeeAdapterMock := new(EmployeeAdapterMock)

	employeeAdapterMock.On("GetAllEmployees").Return([]entities.Employee{}, errors.New("error"))

	newRepository := employee.NewEmployeeRepository(employeeAdapterMock)

	_, err := newRepository.GetEmployeeById(1)

	assert.Error(t, err)
}


