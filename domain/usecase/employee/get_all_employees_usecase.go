package employee

import (
	"rest_app/domain/model/entities/entities"
	"rest_app/domain/model/entities/gateways"
)

type GetAllEmployeesUseCase struct {
	employeeGateway gateways.EmployeeGateway
}

func NewGetAllEmployeesUseCase(employeeGateway gateways.EmployeeGateway) *GetAllEmployeesUseCase {
	return &GetAllEmployeesUseCase{employeeGateway}
}

func (usecase *GetAllEmployeesUseCase) GetAllEmployees() ([]entities.Employee, error) {

	employees, err := usecase.employeeGateway.GetAllEmployees()

	if err != nil {
		return nil, err
	}

	for index := range employees {
		employees[index].CalculateAnnualSalary()
	}

	return employees, nil
}
