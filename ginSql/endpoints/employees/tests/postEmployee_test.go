package employees_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golangLearning/ginSQL/controllers/employees/mocks"
	"github.com/golangLearning/ginSQL/controllers/employees/models"
	"github.com/golangLearning/ginSQL/endpoints/employees"
	"github.com/stretchr/testify/suite"
)

type PostEmployeeTestSuite struct {
	suite.Suite
	controller *mocks.IEmployeeController
	endpoint   employees.PostEmployeeEndpoint
}

func TestPostEmployeeSuite(t *testing.T) {
	suite.Run(t, new(PostEmployeeTestSuite))
}

func (t *PostEmployeeTestSuite) SetupTest() {
	t.controller = new(mocks.IEmployeeController)
	t.endpoint = employees.PostEmployeeEndpoint{Controller: t.controller}
}

func (t *PostEmployeeTestSuite) TestCreateEmployee_Success() {
	writer := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(writer)

	employeeInputObject := models.Employee{
		ID:       1,
		JobTitle: "programmer",
	}

	bodyData, _ := json.Marshal(employeeInputObject)
	bytesBuffer := bytes.NewBuffer(bodyData)
	req := httptest.NewRequest("POST", "/employees", bytesBuffer)
	c.Request = req

	t.controller.On("CreateEmployee", employeeInputObject).Return(employeeInputObject, nil)
	t.endpoint.CreateEmployee(c)

	t.Equal(http.StatusCreated, writer.Code)
	t.Equal(bodyData, writer.Body.Bytes())
	t.controller.AssertNumberOfCalls(t.Suite.T(), "CreateEmployee", 1)
}
