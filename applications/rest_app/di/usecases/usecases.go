package usecases

import (
	"github.com/google/wire"
	"rest_app/domain/usecase/employee"
)

//ProviderSet Provider of Injection Dependencies
var ProviderSet = wire.NewSet(
	employee.NewGetAllEmployeesUseCase,
	employee.NewGetEmployeeByIdUseCase,
)
