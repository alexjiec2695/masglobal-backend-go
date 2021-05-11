package repositories

import (
	"github.com/google/wire"
	"rest_app/domain/model/entities/gateways"
	"rest_app/infrastructure/drivenadapters/repositories/employee"
)

var ProviderSet = wire.NewSet(
	employee.NewEmployeeRepository,
	wire.Bind(
		new(gateways.EmployeeGateway),
		new(*employee.Repository),
	),
)
