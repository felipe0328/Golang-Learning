package employees

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golangLearning/ginSQL/controllers/employees"
	"github.com/golangLearning/ginSQL/controllers/employees/models"
)

type PostEmployeeEndpoint struct {
	controller employees.IEmployeeController
}

func (endpoint *PostEmployeeEndpoint) CreateEmployee(c *gin.Context) {
	var employeeObject models.Employee
	if err := c.ShouldBind(employeeObject); err != nil {
		c.AbortWithError(http.StatusUnprocessableEntity, err)
		return
	}

}
