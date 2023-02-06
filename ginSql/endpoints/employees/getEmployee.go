package employees

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golangLearning/ginSQL/controllers/employees"
	controllerModels "github.com/golangLearning/ginSQL/controllers/employees/models"
	"github.com/golangLearning/ginSQL/endpoints/employees/models"
)

type GetEmployeeEndpoint struct {
	controller employees.IEmployeeController
}

//	@Summary		Get Employee
//	@Description	Get Employee data by id
//	@Produce		json
//	@Param			employeeID	path		integer	true	"employee ID"
//	@Success		200			{object}	controllerModels.Employee
//	@Failure		422			{object}	error
//	@Failure		400			{object}	error
//	@Router			/employees/:id [get]
func (endpoint GetEmployeeEndpoint) GetEmployee(c *gin.Context) {
	var employeeId models.GetEmployeeInput
	if err := c.ShouldBind(&employeeId); err != nil {
		c.AbortWithError(http.StatusUnprocessableEntity, err)
	}

	var employee controllerModels.Employee

	employee, err := endpoint.controller.GetEmployee(employeeId.ID)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, employee)
}
