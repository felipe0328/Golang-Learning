package employees

import (
	"github.com/golangLearning/ginSQL/dals/employees"
	"github.com/golangLearning/ginSQL/dals/employees/models"
)

type IEmployeeController interface {
	GetEmployeesList() ([]models.Employee, error)
	GetEmployee(employeeId int) (models.Employee, error)
	CreateEmployee(employeeData models.Employee) (models.Employee, error)
	UpdateEmployee(employeeID int, employeeData models.Employee) (models.Employee, error)
	DeleteEmployee(employeeId int) error
}

type EmployeeController struct {
	Dal employees.IEmployeeDal
}

func (ec *EmployeeController) GetEmployeesList() ([]models.Employee, error) {
	return ec.Dal.GetEmployeees()
}

func (ec *EmployeeController) GetEmployee(employeeId int) (models.Employee, error) {
	return ec.Dal.GetEmployee(employeeId)
}

func (ec *EmployeeController) CreateEmployee(employeeData models.Employee) (models.Employee, error) {
	return ec.Dal.CreateEmployee(employeeData)
}

func (ec *EmployeeController) UpdateEmployee(employeeID int, employeeData models.Employee) (models.Employee, error) {
	return ec.Dal.UpdateEmployee(employeeID, employeeData)
}

func (ec *EmployeeController) DeleteEmployee(employeeId int) error {
	return ec.Dal.DeleteEmployee(employeeId)
}
