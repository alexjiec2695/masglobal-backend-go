package routers

import (
	"github.com/gin-gonic/gin"
)

//BasePath This const allows define the root path for all rest APIs
const BasePath = "api"

type Router struct {
	R *gin.Engine
}

func NewRouter(r *gin.Engine, employee *EmployeeRouter) *Router {
	employee.registry(r)
	return &Router{r}
}
