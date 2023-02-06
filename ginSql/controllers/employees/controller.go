package employees

import "github.com/golangLearning/ginSQL/controllers/employees/models"

type IEmployeeController interface {
	GetEmployeesList() ([]models.Employee, error)
	GetEmployee(employeeId int) (models.Employee, error)
	CreateEmployee(models.Employee) (models.Employee, error)
	UpdateEmployee(models.Employee) (models.Employee, error)
	DeleteEmployee(employeeId int) error
}
