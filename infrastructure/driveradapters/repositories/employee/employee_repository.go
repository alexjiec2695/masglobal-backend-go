package employee

import (
	"rest_app/domain/model/entities/entities"
	"rest_app/infrastructure/drivenadapters/adapters/employee"
)

type Repository struct {
	employeeAdapter employee.Adapter
}

func NewEmployeeRepository(employeeAdapter employee.Adapter) *Repository {
	return &Repository{employeeAdapter}
}

func (repository *Repository) GetAllEmployees() ([]entities.Employee, error) {
	return repository.employeeAdapter.GetAllEmployees()
}

func (repository *Repository) GetEmployeeById(ID int) (entities.Employee, error) {
	employees, err := repository.employeeAdapter.GetAllEmployees()

	if err != nil {
		return entities.Employee{}, err
	}

	for _, item := range employees {
		if item.Id == ID {
			return item, nil
		}
	}

	return entities.Employee{}, nil
}
