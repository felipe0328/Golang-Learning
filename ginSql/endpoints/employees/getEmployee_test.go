package employees

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golangLearning/ginSQL/controllers/employees/mocks"
	controllerModels "github.com/golangLearning/ginSQL/controllers/employees/models"
	"github.com/golangLearning/ginSQL/endpoints/employees/models"
	"github.com/stretchr/testify/suite"
)

type GetEmployeeTestSuite struct {
	suite.Suite
	controller *mocks.IEmployeeController
	endpoint   GetEmployeeEndpoint
}

func EmployeeGetTest(t *testing.T) {
	suite.Run(t, new(GetEmployeeTestSuite))
}

func (t *GetEmployeeTestSuite) SetupTest() {
	t.controller = new(mocks.IEmployeeController)
	t.endpoint = GetEmployeeEndpoint{controller: t.controller}
}

func (t *GetEmployeeTestSuite) TestGetEmployee_Success() {
	testUserID := 1

	writer := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(writer)
	c.Set("id", testUserID)

	userEmployee := models.GetEmployeeInput{ID: testUserID}
	employeeResult := controllerModels.Employee{ID: testUserID, Phone: "1234"}
	t.controller.On("GetEmployee", userEmployee).Return(employeeResult, nil)

	expectedResult, _ := json.Marshal(employeeResult)

	t.endpoint.GetEmployee(c)

	t.Equal(http.StatusOK, writer.Code)
	t.Equal(expectedResult, writer.Body.Bytes())
	t.controller.AssertNumberOfCalls(t.Suite.T(), "GetEmployee", 1)

	var result controllerModels.Employee
	json.Unmarshal(writer.Body.Bytes(), &result)
	t.Equal(testUserID, result.ID)
	t.Equal("1234", result.Phone)
}

func (t *GetEmployeeTestSuite) TestGetEmployee_NoParams_Failed() {
	writer := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(writer)

	t.endpoint.GetEmployee(c)

	t.Equal(http.StatusUnprocessableEntity, writer.Code)
	t.NotNil(c.Errors.Last().Err)
	t.controller.AssertNumberOfCalls(t.Suite.T(), "GetEmployee", 0)
}

func (t *GetEmployeeTestSuite) TestGetEmployee_ControllerError_Failed() {
	testUserID := 1

	writer := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(writer)
	c.Set("id", testUserID)

	userEmployee := models.GetEmployeeInput{ID: testUserID}
	expectedError := errors.New("expected error.")
	employeeResult := controllerModels.Employee{}
	t.controller.On("GetEmployee", userEmployee).Return(employeeResult, expectedError)

	t.endpoint.GetEmployee(c)

	t.Equal(http.StatusBadRequest, writer.Code)
	t.controller.AssertNumberOfCalls(t.Suite.T(), "GetEmployee", 1)
	t.Equal(c.Errors.Last().Err, expectedError)
}
