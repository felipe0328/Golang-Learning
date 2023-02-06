package employees

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golangLearning/ginSQL/controllers/employees"
	"github.com/golangLearning/ginSQL/controllers/employees/models"
	"github.com/golangLearning/ginSQL/endpoints/employees/utils"
)

type PostEmployeeEndpoint struct {
	Controller employees.IEmployeeController
}

// @Summary		Create Employee
// @Tags			Employees
// @Description	Create employee in DB
// @Produce		json
// @Param			body	body		models.Employee	true	"Employee Object"
// @Success		200		{object}	models.Employee
// @Failure		400		{object}	error
// @Failure		422		{object}	error
// @Router			/employees [post]
func (endpoint *PostEmployeeEndpoint) CreateEmployee(c *gin.Context) {
	var employeeObject models.Employee
	if err := c.ShouldBindJSON(&employeeObject); err != nil {
		c.AbortWithError(http.StatusUnprocessableEntity, utils.ErrMissingEmployeeData)
		return
	}

	resultEmployee, err := endpoint.Controller.CreateEmployee(employeeObject)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, resultEmployee)
}
