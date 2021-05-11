package employee

import "rest_app/domain/model/entities/entities"

type Adapter interface {
	GetAllEmployees() ([]entities.Employee, error)
}