package routers

import (
	"github.com/gin-gonic/gin"
	"rest_app/infrastructure/entrypoints/rest/handlers"
)

//EmployeeRouter This router allows register all endpoints associated with Employee
type EmployeeRouter struct {
	handler *handlers.EmployeeHandler
}

//NewEmployeeRouter this function is the EmployeeRouter constructor instance
func NewEmployeeRouter(handler *handlers.EmployeeHandler) *EmployeeRouter {
	ha := &EmployeeRouter{handler}
	return ha
}

func (h *EmployeeRouter) registry(router *gin.Engine) {

	base := BasePath
	group := router.Group(base)

	group.GET("employees", h.handler.Employee)
	group.GET("employees/:id", h.handler.EmployeeById)

}