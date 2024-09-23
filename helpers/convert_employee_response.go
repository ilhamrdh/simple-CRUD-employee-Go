package helpers

import (
	"github.com/ilhamrdh/simple-CRUD-employee-Go.git/models/entity"
	"github.com/ilhamrdh/simple-CRUD-employee-Go.git/models/web"
)

func ConvertEmployeeResponse(employee entity.Employees) web.EmployeeResponse {
	return web.EmployeeResponse{
		EmployeeId: employee.EmployeeId,
		Fullname:   employee.Fullname,
		Address:    employee.Addrees,
	}
}

func ConvertEmployeeResponses(employees []entity.Employees) []web.EmployeeResponse {
	var employeeResponses []web.EmployeeResponse
	for _, employee := range employees {
		employeeResponses = append(employeeResponses, ConvertEmployeeResponse(employee))
	}
	return employeeResponses
}
