package employees

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golangLearning/ginSQL/controllers/employees"
	controllerModels "github.com/golangLearning/ginSQL/dals/employees/models"
	"github.com/golangLearning/ginSQL/endpoints/employees/models"
	"github.com/golangLearning/ginSQL/endpoints/employees/utils"
)

type GetEmployeeEndpoint struct {
	Controller employees.IEmployeeController
}

// @Summary		Get Employee
// @Tags			Employee
// @Description	Get Employee data by id
// @Produce		json
// @Param			employeeID	path		integer	true	"employee ID"
// @Success		200			{object}	controllerModels.Employee
// @Failure		422			{object}	error
// @Failure		400			{object}	error
// @Router			/employees/:id [get]
func (endpoint GetEmployeeEndpoint) GetEmployee(c *gin.Context) {
	var employeeInput models.GetEmployeeInput
	err := c.BindUri(&employeeInput)
	if err != nil {
		c.AbortWithError(http.StatusUnprocessableEntity, utils.ErrMissingEmployeeId)
	}

	var employee controllerModels.Employee

	employee, err = endpoint.Controller.GetEmployee(employeeInput.ID)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, employee)
}
