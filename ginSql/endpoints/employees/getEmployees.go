package employees

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golangLearning/ginSQL/controllers/employees"
)

type GetEmployeesEndpoint struct {
	controller employees.IEmployeeController
}

// getEmployees godoc
//
//	@Summary		Get Employees
//	@Description	Get List of employees
//	@Produce		json
//	@Success		200	{object}	[]models.Employee
//	@Failure		400	{object}	error
//	@Router			/employees [get]
func (endpoint GetEmployeesEndpoint) GetEmployees(c *gin.Context) {
	responseEmployees, err := endpoint.controller.GetEmployeesList()

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, responseEmployees)
}
