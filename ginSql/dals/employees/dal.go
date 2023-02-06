package employees

import "github.com/golangLearning/ginSQL/dals/employees/models"

type IEmployeeDal interface {
	GetEmployeees() ([]models.Employee, error)
	GetEmployee(employeeID int) (models.Employee, error)
	CreateEmployee(employee models.Employee) (models.Employee, error)
	UpdateEmployee(employeeID int, employee models.Employee) (models.Employee, error)
	DeleteEmployee(employee int) error
}
