package db

import (
	"errors"

	"github.com/go-rest/models"
)

type EmployeeCollection struct {
	Employees []models.Employee
}

var empCol *EmployeeCollection

// to make sure only one instance is returned
func GetEmployeeInstance() *EmployeeCollection {
	if collection == nil {
		empCol = &EmployeeCollection{}
	}

	return empCol
}

func (e *EmployeeCollection) NextId() int {
	return len(e.Employees) + 1
}

func (e *EmployeeCollection) List() []models.Employee {
	return e.Employees
}

func (e *EmployeeCollection) Create(employee models.Employee) string {
	e.Employees = append(e.Employees, employee)
	return "Employee added successfully"
}

func (e *EmployeeCollection) Update(employee models.Employee) (error, string) {
	found := false

	for indx, value := range e.Employees {
		if value.Id == employee.Id {
			e.Employees[indx] = employee
			found = true
			break
		}
	}

	if found {
		return nil, "Updated successfully"
	}

	return errors.New("Not Found"), ""
}

func (e *EmployeeCollection) FindById(employeeId int) (error, models.Employee) {

	for _, value := range e.Employees {
		if value.Id == employeeId {
			return nil, value
		}
	}

	return errors.New("Not Found"), models.Employee{}
}

func (e *EmployeeCollection) Delete(employeeId int) (error, string) {
	removeIndx := -1

	for indx, value := range e.Employees {
		if value.Id == employeeId {
			removeIndx = indx
			break
		}
	}

	if removeIndx != -1 {
		e.Employees = append(e.Employees[:removeIndx], e.Employees[removeIndx+1:]...)
		return nil, "Deleted"
	}

	return errors.New("Not Found"), ""
}
