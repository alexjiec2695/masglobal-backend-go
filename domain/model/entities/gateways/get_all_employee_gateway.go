package gateways

import "rest_app/domain/model/entities/entities"

type EmployeeGateway interface {
	GetAllEmployees() ([]entities.Employee, error)
	GetEmployeeById(ID int) (entities.Employee, error)
}
