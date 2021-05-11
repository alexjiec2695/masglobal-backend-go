package externalservice

import (
	"github.com/google/wire"
	employeeAdapter "rest_app/infrastructure/drivenadapters/adapters/employee"
	"rest_app/infrastructure/drivenadapters/externalservice/employee"
)

var ProviderSet = wire.NewSet(
	employee.NewExternalServiceImpl,
	wire.Bind(
		new(employeeAdapter.Adapter),
		new(*employee.ExternalServiceImpl),
	),
)
