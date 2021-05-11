package employee

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"rest_app/domain/model/entities/entities"
	"rest_app/infrastructure/drivenadapters/externalservice/entitydata"
)

type ExternalServiceImpl struct {
}

func NewExternalServiceImpl() *ExternalServiceImpl {
	return &ExternalServiceImpl{}
}

const Baseurl = "http://masglobaltestapi.azurewebsites.net/api/employees"

func (e *ExternalServiceImpl) GetAllEmployees() ([]entities.Employee, error) {

	var employeeData []entitydata.EmployeeData

	resp, err := http.Get(Baseurl)

	if err != nil || resp.StatusCode > 399 {
		return []entities.Employee{}, errors.New("gateway error")
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	json.Unmarshal(body, &employeeData)

	employees := make([]entities.Employee, len(employeeData))

	for index, item := range employeeData {
		employees[index] = entities.Employee{
			Id:               item.Id,
			Name:             item.Name,
			ContractTypeName: item.ContractTypeName,
			HourlySalary:     item.HourlySalary,
			MonthlySalary:    item.MonthlySalary,
			RoleId:           item.RoleId,
			RoleDescription:  item.RoleDescription,
			RoleName:         item.RoleName,
		}
	}

	return employees, nil
}

