package employees

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golangLearning/ginSQL/controllers/employees"
	"github.com/golangLearning/ginSQL/endpoints/employees/models"
	"github.com/golangLearning/ginSQL/endpoints/employees/utils"
)

type DeleteEmployeeEndpoint struct {
	Controller employees.IEmployeeController
}

// @Summary Delete Employee
// @Tags Employees
// @Description Delete Employee from DB
// @Param employeeID path int true "Employee ID"
// @Success 204
// @Failure 400 {object} error
// @Router /employees/:id [delete]
func (endpoint *DeleteEmployeeEndpoint) DeleteEmployee(c *gin.Context) {
	var employeeID models.GetEmployeeInput

	if err := c.ShouldBindUri(&employeeID); err != nil {
		c.AbortWithError(http.StatusBadRequest, utils.ErrMissingEmployeeId)
		return
	}

	err := endpoint.Controller.DeleteEmployee(employeeID.ID)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.AbortWithStatus(http.StatusNoContent)
}
