package employees

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golangLearning/ginSQL/endpoints/employees/models"
)

type GetEmployeesEndpoint struct{}

// getEmployees godoc
// @Summary Get Employees
// @Description Get List of employees
// @Produce json
// @Success 200 {object} []models.Employee
// @Failure 400
// @Router /employees [get]
func (endpoint GetEmployeesEndpoint) GetEmployees(c *gin.Context) {
	responseEmployees := models.Employee{}
	c.JSON(http.StatusOK, responseEmployees)
}
