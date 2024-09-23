package services

import (
	"context"

	"github.com/ilhamrdh/simple-CRUD-employee-Go.git/models/web"
)

type EmployeeService interface {
	Create(ctx context.Context, request web.EmployeeCreateRequest) web.EmployeeResponse
	Update(ctx context.Context, request web.EmployeeUpdateRequest) web.EmployeeResponse
	Delete(ctx context.Context, employeeId string)
	GetById(ctx context.Context, employeeId string) web.EmployeeResponse
	GetAll(ctx context.Context) []web.EmployeeResponse
}
