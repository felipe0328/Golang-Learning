package employees

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golangLearning/ginSQL/controllers/employees"
	"github.com/golangLearning/ginSQL/controllers/employees/models"
	endpointModels "github.com/golangLearning/ginSQL/endpoints/employees/models"
	"github.com/golangLearning/ginSQL/endpoints/employees/utils"
)

type PatchEmployeeEndpoint struct {
	Controller employees.IEmployeeController
}

// @Summary		Patch Employee
// @Tags			Employees
// @Description	Update Employee
// @Produce		json
// @Param			employeeID	path		integer			true	"Employee ID"
// @Param			body		body		models.Employee	true	"Employee Object"
// @Success		202			{object}	models.Employee
// @Failure		422			{object}	error
// @Failure		400			{object}	error
// @Router			/employees/:id [patch]
func (endpoint *PatchEmployeeEndpoint) UpdateEmployee(c *gin.Context) {
	var employeeID endpointModels.GetEmployeeInput
	if err := c.ShouldBindUri(&employeeID); err != nil {
		fmt.Println("Returnin with 422")
		c.AbortWithError(http.StatusUnprocessableEntity, utils.ErrMissingEmployeeId)
		return
	}

	var employeeObject models.Employee
	if err := c.ShouldBindJSON(&employeeObject); err != nil {
		c.AbortWithError(http.StatusUnprocessableEntity, utils.ErrMissingEmployeeData)
		return
	}

	employeeUpdated, err := endpoint.Controller.UpdateEmployee(employeeID.ID, employeeObject)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusAccepted, employeeUpdated)
}
