package employees

import "github.com/golangLearning/ginSQL/controllers/employees/models"

type IEmployeeController interface {
	GetEmployeesList() ([]models.Employee, error)
}
