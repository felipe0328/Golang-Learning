package employees_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golangLearning/ginSQL/controllers/employees/mocks"
	"github.com/golangLearning/ginSQL/controllers/employees/models"
	"github.com/golangLearning/ginSQL/endpoints/employees"
	"github.com/stretchr/testify/suite"
)

type GetEmployeesTestSuite struct {
	suite.Suite
	controller *mocks.IEmployeeController
	endpoint   employees.GetEmployeesEndpoint
}

func (t *GetEmployeesTestSuite) SetupTest() {
	t.controller = new(mocks.IEmployeeController)
	t.endpoint = employees.GetEmployeesEndpoint{Controller: t.controller}
}

func TestGetEndpointSuite(t *testing.T) {
	suite.Run(t, new(GetEmployeesTestSuite))
}

func (t *GetEmployeesTestSuite) TestGetEndpoint_Success() {
	controllerOutput := make([]models.Employee, 0)

	expectedOutput, _ := json.Marshal(controllerOutput)
	t.controller.On("GetEmployeesList").Return(controllerOutput, nil)

	writer := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(writer)

	t.endpoint.GetEmployees(c)

	t.Equal(http.StatusOK, writer.Code)
	t.Equal(expectedOutput, writer.Body.Bytes())
	t.controller.AssertNumberOfCalls(t.Suite.T(), "GetEmployeesList", 1)
}

func (t *GetEmployeesTestSuite) TestGetEndpoint_ControllerError_Failed() {
	expectedError := errors.New("expected error")

	t.controller.On("GetEmployeesList").Return(nil, expectedError)

	writer := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(writer)

	t.endpoint.GetEmployees(c)

	t.Equal(http.StatusBadRequest, writer.Code)
	t.Equal(expectedError, c.Errors.Last().Err)
	t.controller.AssertNumberOfCalls(t.Suite.T(), "GetEmployeesList", 1)
}
