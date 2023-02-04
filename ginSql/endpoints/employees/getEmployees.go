package employees

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golangLearning/ginSQL/endpoints/employees/models"
)

// getEmployees godoc
// @Summary Get Employees
// @Description Get List of employees
// @Produce json
// @Success 200 {object} []models.Employee
// @Failure 400 {string} string "err_code：10002"
// @Router /employees [get]
func GetEmployees(c *gin.Context) {
	responseEmployees := models.Employee{}
	c.JSON(http.StatusOK, responseEmployees)
}
