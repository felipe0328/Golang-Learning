package employees

import "github.com/gin-gonic/gin"

func Routes(r *gin.Engine) {
	////////// Implementations
	employeesGet := GetEmployeesEndpoint{}
	employeeGet := GetEmployeeEndpoint{}
	employeePost := PostEmployeeEndpoint{}
	employeePatch := PatchEmployeeEndpoint{}
	employeeDelete := DeleteEmployeeEndpoint{}

	///////// Endpoints
	r.GET("/employees", employeesGet.GetEmployees)
	r.POST("/employees", employeePost.CreateEmployee)
	r.GET("/employees/:id", employeeGet.GetEmployee)
	r.PATCH("/employees/:id", employeePatch.UpdateEmployee)
	r.DELETE("/employees/:id", employeeDelete.DeleteEmployee)
}
