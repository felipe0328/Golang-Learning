package employees

import (
	"github.com/gin-gonic/gin"
	"github.com/golangLearning/ginSQL/controllers/employees"
	employeesDal "github.com/golangLearning/ginSQL/dals/employees"
	"github.com/jackc/pgx"
)

func Routes(r *gin.Engine, db *pgx.ConnPool) {
	////////// Dal
	employeesDal := &employeesDal.EmployeeDal{ConnPool: db}

	////////// Controller
	employeesController := &employees.EmployeeController{Dal: employeesDal}

	////////// Implementations
	employeesGet := GetEmployeesEndpoint{Controller: employeesController}
	employeeGet := GetEmployeeEndpoint{Controller: employeesController}
	employeePost := PostEmployeeEndpoint{Controller: employeesController}
	employeePatch := PatchEmployeeEndpoint{Controller: employeesController}
	employeeDelete := DeleteEmployeeEndpoint{Controller: employeesController}

	///////// Endpoints
	r.GET("/employees", employeesGet.GetEmployees)
	r.POST("/employees", employeePost.CreateEmployee)
	r.GET("/employees/:id", employeeGet.GetEmployee)
	r.PATCH("/employees/:id", employeePatch.UpdateEmployee)
	r.DELETE("/employees/:id", employeeDelete.DeleteEmployee)
}
