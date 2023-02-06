package employees_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golangLearning/ginSQL/controllers/employees/mocks"
	"github.com/golangLearning/ginSQL/endpoints/employees"
	"github.com/golangLearning/ginSQL/endpoints/employees/utils"
	"github.com/stretchr/testify/suite"
)

type DeleteEmployeeSuite struct {
	suite.Suite
	controller *mocks.IEmployeeController
	endpoint   employees.DeleteEmployeeEndpoint
}

func TestDeleteEmployee(t *testing.T) {
	suite.Run(t, new(DeleteEmployeeSuite))
}

func (t *DeleteEmployeeSuite) SetupTest() {
	t.controller = new(mocks.IEmployeeController)
	t.endpoint = employees.DeleteEmployeeEndpoint{Controller: t.controller}
}

func (t *DeleteEmployeeSuite) TestDeleteEmployee_Success() {
	testUserID := "1"

	writer := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(writer)
	c.AddParam("id", testUserID)

	t.controller.On("DeleteEmployee", 1).Return(nil)
	t.endpoint.DeleteEmployee(c)

	t.Equal(http.StatusNoContent, writer.Code)
	t.controller.AssertNumberOfCalls(t.T(), "DeleteEmployee", 1)
}

func (t *DeleteEmployeeSuite) TestDeleteEmployee_MissingID_Failed() {
	writer := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(writer)

	t.endpoint.DeleteEmployee(c)

	t.Equal(http.StatusBadRequest, writer.Code)
	t.NotNil(c.Errors.Last().Err)
	t.Equal(utils.ErrMissingEmployeeId, c.Errors.Last().Err)
	t.controller.AssertNumberOfCalls(t.T(), "DeleteEmployee", 0)
}

func (t *DeleteEmployeeSuite) TestDeleteEmployee_ControllerError_Failed() {
	testUserID := "1"

	writer := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(writer)
	c.AddParam("id", testUserID)

	expectedError := errors.New("expected error")
	t.controller.On("DeleteEmployee", 1).Return(expectedError)
	t.endpoint.DeleteEmployee(c)

	t.Equal(http.StatusBadRequest, writer.Code)
	t.NotNil(c.Errors.Last().Err)
	t.Equal(expectedError, c.Errors.Last().Err)
	t.controller.AssertNumberOfCalls(t.T(), "DeleteEmployee", 1)
}
