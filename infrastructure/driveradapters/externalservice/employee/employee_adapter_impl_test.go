package employee_test

import (
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"rest_app/infrastructure/drivenadapters/externalservice/employee"
	"rest_app/infrastructure/drivenadapters/externalservice/entitydata"
	"testing"
)

func TestExternalServiceImpl_GetAllEmployees_OK(t *testing.T) {

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	employeesData := make([]entitydata.EmployeeData, 1)
	employeesData[0] = entitydata.EmployeeData{
		Id:               1,
		Name:             "test",
		ContractTypeName: "HourlySalaryEmployee",
		RoleId:           1,
		RoleName:         "test",
		RoleDescription:  "",
		HourlySalary:     0,
		MonthlySalary:    0,
	}

	responder, err := httpmock.NewJsonResponder(200, employeesData)

	httpmock.RegisterResponder("GET", employee.Baseurl, responder)

	externalService := employee.NewExternalServiceImpl()

	employeeResult, err := externalService.GetAllEmployees()

	assert.NoError(t, err)
	assert.Equal(t, employeesData[0].Id, employeeResult[0].Id)
}

func TestExternalServiceImpl_GetAllEmployees_Error_Rest(t *testing.T) {

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	responder, err := httpmock.NewJsonResponder(500, map[string]interface{}{})

	httpmock.RegisterResponder("GET", employee.Baseurl, responder)

	externalService := employee.NewExternalServiceImpl()

	_, err = externalService.GetAllEmployees()

	assert.Error(t, err)
}

func TestExternalServiceImpl_GetEmployeeById(t *testing.T) {

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	employeesData := make([]entitydata.EmployeeData, 1)
	employeesData[0] = entitydata.EmployeeData{
		Id:               1,
		Name:             "test",
		ContractTypeName: "HourlySalaryEmployee",
		RoleId:           1,
		RoleName:         "test",
		RoleDescription:  "",
		HourlySalary:     0,
		MonthlySalary:    0,
	}

	responder, err := httpmock.NewJsonResponder(200, employeesData)

	httpmock.RegisterResponder("GET", employee.Baseurl, responder)

	externalService := employee.NewExternalServiceImpl()

	employeeResult, err := externalService.GetAllEmployees()

	assert.NoError(t, err)
	assert.Equal(t, employeesData[0].Id, employeeResult[0].Id)
}
