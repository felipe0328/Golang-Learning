package employees

import (
	"github.com/golangLearning/ginSQL/dals/employees/models"
	"github.com/jackc/pgx"
)

type IEmployeeDal interface {
	GetEmployeees() ([]models.Employee, error)
	GetEmployee(employeeID int) (models.Employee, error)
	CreateEmployee(employee models.Employee) (models.Employee, error)
	UpdateEmployee(employeeID int, employee models.Employee) (models.Employee, error)
	DeleteEmployee(employee int) error
}

type EmployeeDal struct {
	*pgx.ConnPool
}

func (d *EmployeeDal) GetEmployeees() ([]models.Employee, error) {
	query := `Select * from employees`
	rows, err := d.Query(query)

	employeesData := make([]models.Employee, 0)

	if err != nil {
		return employeesData, err
	}

	for rows.Next() {
		employee := models.Employee{}
		err := rows.Scan(
			&employee.ID,
			&employee.City,
			&employee.Country,
			&employee.Email,
			&employee.FirstName,
			&employee.FullName,
			&employee.Image,
			&employee.JobTitle,
			&employee.LastName,
			&employee.OnboardingCompletion,
			&employee.Phone,
		)

		if err != nil {
			return employeesData, err
		}

		employeesData = append(employeesData, employee)
	}

	return employeesData, nil
}

func (d *EmployeeDal) GetEmployee(employeeID int) (models.Employee, error) {
	employee := models.Employee{}
	return employee, nil
}

func (d *EmployeeDal) CreateEmployee(employee models.Employee) (models.Employee, error) {
	newEmployee := models.Employee{}
	return newEmployee, nil
}

func (d *EmployeeDal) UpdateEmployee(employeeID int, employee models.Employee) (models.Employee, error) {
	updatedEmployee := models.Employee{}
	return updatedEmployee, nil
}

func (d *EmployeeDal) DeleteEmployee(employee int) error {
	return nil
}
