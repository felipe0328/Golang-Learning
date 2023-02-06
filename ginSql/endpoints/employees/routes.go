package employees

import "github.com/gin-gonic/gin"

func Routes(r *gin.Engine) {
	////////// Implementations
	employeesGet := GetEmployeesEndpoint{}
	employeeGet := GetEmployeeEndpoint{}

	///////// Endpoints
	r.GET("/employees", employeesGet.GetEmployees)
	r.POST("/employees")
	r.GET("/employees/:id", employeeGet.GetEmployee)
	r.PATCH("/employees/:id")
	r.DELETE("/employees/:id")
}
