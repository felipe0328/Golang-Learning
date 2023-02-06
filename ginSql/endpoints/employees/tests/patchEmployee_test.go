package employees_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golangLearning/ginSQL/controllers/employees/mocks"
	"github.com/golangLearning/ginSQL/dals/employees/models"
	"github.com/golangLearning/ginSQL/endpoints/employees"
	"github.com/golangLearning/ginSQL/endpoints/employees/utils"
	"github.com/stretchr/testify/suite"
)

type PatchEmployeeTestSuite struct {
	suite.Suite
	controller *mocks.IEmployeeController
	endpoint   employees.PatchEmployeeEndpoint
}

func TestPatchEmployees(t *testing.T) {
	suite.Run(t, new(PatchEmployeeTestSuite))
}

func (t *PatchEmployeeTestSuite) SetupTest() {
	t.controller = new(mocks.IEmployeeController)
	t.endpoint = employees.PatchEmployeeEndpoint{Controller: t.controller}
}

func (t *PatchEmployeeTestSuite) TestUpdateEmployee_Success() {
	writer := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(writer)

	testUserID := "1"
	employeeInputObject := models.Employee{
		JobTitle: "programmer",
	}

	bodyData, _ := json.Marshal(employeeInputObject)
	bytesBuffer := bytes.NewBuffer(bodyData)
	req := httptest.NewRequest("POST", "/employees", bytesBuffer)
	c.Request = req
	c.AddParam("id", testUserID)

	returnedEmployeeObject := models.Employee{
		JobTitle:  "Programmer",
		FirstName: "The",
		LastName:  "Programmer",
	}
	t.controller.On("UpdateEmployee", 1, employeeInputObject).Return(returnedEmployeeObject, nil)
	t.endpoint.UpdateEmployee(c)

	expectedResult, _ := json.Marshal(returnedEmployeeObject)

	t.Equal(http.StatusAccepted, writer.Code)
	t.Equal(expectedResult, writer.Body.Bytes())
	t.controller.AssertNumberOfCalls(t.T(), "UpdateEmployee", 1)
}

func (t *PatchEmployeeTestSuite) TestUpdateEmployee_NoCorrectPathParams_Failed() {
	writer := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(writer)

	employeeInputObject := models.Employee{
		JobTitle: "programmer",
	}

	bodyData, _ := json.Marshal(employeeInputObject)
	bytesBuffer := bytes.NewBuffer(bodyData)
	req := httptest.NewRequest("POST", "/employees", bytesBuffer)
	c.Request = req

	t.endpoint.UpdateEmployee(c)

	t.Equal(http.StatusUnprocessableEntity, writer.Code)
	t.NotNil(c.Errors.Last().Err)
	t.Equal(utils.ErrMissingEmployeeId, c.Errors.Last().Err)
	t.controller.AssertNumberOfCalls(t.T(), "UpdateEmployee", 0)
}

func (t *PatchEmployeeTestSuite) TestUpdateEmployee_NoBodyParam_Failed() {
	writer := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(writer)

	testUserID := "1"

	req := httptest.NewRequest("POST", "/employees", nil)
	c.Request = req
	c.AddParam("id", testUserID)

	t.endpoint.UpdateEmployee(c)

	t.Equal(http.StatusUnprocessableEntity, writer.Code)
	t.NotNil(c.Errors.Last().Err)
	t.Equal(utils.ErrMissingEmployeeData, c.Errors.Last().Err)
	t.controller.AssertNumberOfCalls(t.T(), "UpdateEmployee", 0)
}

func (t *PatchEmployeeTestSuite) TestUpdateEmployee_ControllerError_Failed() {
	writer := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(writer)

	testUserID := "1"
	employeeInputObject := models.Employee{
		JobTitle: "programmer",
	}

	bodyData, _ := json.Marshal(employeeInputObject)
	bytesBuffer := bytes.NewBuffer(bodyData)
	req := httptest.NewRequest("POST", "/employees", bytesBuffer)
	c.Request = req
	c.AddParam("id", testUserID)

	expectedError := errors.New("expected Error")
	emptyEmployee := models.Employee{}
	t.controller.On("UpdateEmployee", 1, employeeInputObject).Return(emptyEmployee, expectedError)
	t.endpoint.UpdateEmployee(c)

	t.Equal(http.StatusBadRequest, writer.Code)
	t.NotNil(c.Errors.Last().Err)
	t.Equal(expectedError, c.Errors.Last().Err)
	t.controller.AssertNumberOfCalls(t.T(), "UpdateEmployee", 1)
}
