package employees

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golangLearning/ginSQL/endpoints/employees/models"
	"github.com/stretchr/testify/suite"
)

type GetEmployeesTestSuite struct {
	suite.Suite
}

func (t *GetEmployeesTestSuite) SetupTest() {}

func TestGetEndpointSuite(t *testing.T) {
	suite.Run(t, new(GetEmployeesTestSuite))
}

func (t *GetEmployeesTestSuite) TestGetEndpoint() {
	r := gin.Default()
	Routes(r)

	writer := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/employees", nil)
	r.ServeHTTP(writer, req)

	expectedError := make([]models.Employee, 0)

	var result []models.Employee
	unmarshalError := json.Unmarshal(writer.Body.Bytes(), &result)

	t.Equal(http.StatusOK, writer.Code)
	t.Nil(unmarshalError)
	t.Equal(expectedError, result)
}
