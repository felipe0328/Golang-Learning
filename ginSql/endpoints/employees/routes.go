package employees

import "github.com/gin-gonic/gin"

func Routes(r *gin.Engine) {
	////////// Implementations
	employeesGet := GetEmployeesEndpoint{}

	///////// Endpoints
	r.GET("/employees", employeesGet.GetEmployees)
	r.POST("/employees")
	r.GET("/employees/:id")
	r.PATCH("/employees/:id")
	r.PUT("/employees/:id")
	r.DELETE("/employees/:id")
}
