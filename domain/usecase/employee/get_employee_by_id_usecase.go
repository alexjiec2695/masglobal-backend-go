package employee

import (
	"rest_app/domain/model/entities/entities"
	"rest_app/domain/model/entities/gateways"
)

type GetEmployeeByIdUseCase struct {
	employeeGateway gateways.EmployeeGateway
}

func NewGetEmployeeByIdUseCase(employeeGateway gateways.EmployeeGateway) *GetEmployeeByIdUseCase {
	return &GetEmployeeByIdUseCase{employeeGateway}
}

func (usecase *GetEmployeeByIdUseCase) GetEmployeeById(ID int) (entities.Employee, error) {

	employee, err := usecase.employeeGateway.GetEmployeeById(ID)

	if err != nil {
		return entities.Employee{}, err
	}

	employee.CalculateAnnualSalary()

	return employee, nil
}
