package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest_app/domain/usecase/employee"
	"rest_app/infrastructure/entrypoints/rest/dtos/rs"
	"strconv"
)

// EmployeeHandler s
type EmployeeHandler struct {
	getAllEmployeesUseCase *employee.GetAllEmployeesUseCase
	getEmployeeByIdUseCase *employee.GetEmployeeByIdUseCase
}

// NewEmployeeHandler s
func NewEmployeeHandler(getAllEmployeesUseCase *employee.GetAllEmployeesUseCase, getEmployeeByIdUseCase *employee.GetEmployeeByIdUseCase) (employeeHandler *EmployeeHandler) {
	employeeHandler = &EmployeeHandler{getAllEmployeesUseCase, getEmployeeByIdUseCase}
	return
}

func (e *EmployeeHandler) Employee(ctx *gin.Context) {


	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")
	ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
	ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	ctx.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
	ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

	employees, err := e.getAllEmployeesUseCase.GetAllEmployees()

	if err != nil {
		panic(err)
	}

	response := make([]rs.EmployeesRsDto, len(employees))

	for index, item := range employees {
		response[index] = rs.EmployeesRsDto{
			Id:               item.Id,
			RoleName:         item.RoleName,
			RoleDescription:  item.RoleDescription,
			RoleId:           item.RoleId,
			MonthlySalary:    item.MonthlySalary,
			HourlySalary:     item.HourlySalary,
			ContractTypeName: item.ContractTypeName,
			Name:             item.Name,
			AnnualSalary:     item.AnnualSalary,
		}
	}

	ctx.JSON(http.StatusOK, response)
}

func (e *EmployeeHandler) EmployeeById(ctx *gin.Context) {

	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")
	ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
	ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	ctx.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
	ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

	ID, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		panic(err)
	}

	employee, err := e.getEmployeeByIdUseCase.GetEmployeeById(ID)

	if err != nil {
		panic(err)
	}

	response := make([]rs.EmployeesRsDto, 1)

	response[0] = rs.EmployeesRsDto{
		Id:               employee.Id,
		RoleName:         employee.RoleName,
		MonthlySalary:    employee.MonthlySalary,
		HourlySalary:     employee.HourlySalary,
		ContractTypeName: employee.ContractTypeName,
		Name:             employee.Name,
		RoleId:           employee.RoleId,
		RoleDescription:  employee.RoleDescription,
		AnnualSalary:     employee.AnnualSalary,
	}

	ctx.JSON(http.StatusOK, response)

}
