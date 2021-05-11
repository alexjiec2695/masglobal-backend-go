module rest_app/infrastructure/entrypoints/rest

go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/stretchr/testify v1.4.0
	rest_app/domain/model v0.0.0
	rest_app/domain/usecase v0.0.0
)

replace (
	rest_app/domain/model => ../../../domain/model
	rest_app/domain/usecase => ../../../domain/usecase
)
