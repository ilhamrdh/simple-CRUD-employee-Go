package services

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/ilhamrdh/simple-CRUD-employee-Go.git/exception"
	"github.com/ilhamrdh/simple-CRUD-employee-Go.git/helpers"
	"github.com/ilhamrdh/simple-CRUD-employee-Go.git/models/entity"
	"github.com/ilhamrdh/simple-CRUD-employee-Go.git/models/web"
	"github.com/ilhamrdh/simple-CRUD-employee-Go.git/repositories"
)

type EmployeeServiceImpl struct {
	EmployeeRepository repositories.EmployeeRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewEmployeeService(employeeRepository repositories.EmployeeRepository, DB *sql.DB, validate *validator.Validate) EmployeeService {
	return &EmployeeServiceImpl{
		EmployeeRepository: employeeRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

// Create implements EmployeeService.
func (e *EmployeeServiceImpl) Create(ctx context.Context, request web.EmployeeCreateRequest) web.EmployeeResponse {
	err := e.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx, err := e.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	employee := entity.Employees{
		EmployeeId: request.EmployeeId,
		Fullname:   request.Fullname,
		Addrees:    request.Address,
	}

	employee = e.EmployeeRepository.Save(ctx, tx, employee)

	return helpers.ConvertEmployeeResponse(employee)
}

// Delete implements EmployeeService.
func (e *EmployeeServiceImpl) Delete(ctx context.Context, employeeId string) {
	tx, err := e.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	employee, err := e.EmployeeRepository.GetById(ctx, tx, employeeId)
	if err != nil {
		panic(exception.NewNotFoundException(err.Error()))
	}

	e.EmployeeRepository.Delete(ctx, tx, employee)

}

// GetAll implements EmployeeService.
func (e *EmployeeServiceImpl) GetAll(ctx context.Context) []web.EmployeeResponse {
	tx, err := e.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	employees := e.EmployeeRepository.GetAll(ctx, tx)
	return helpers.ConvertEmployeeResponses(employees)
}

// GetById implements EmployeeService.
func (e *EmployeeServiceImpl) GetById(ctx context.Context, employeeId string) web.EmployeeResponse {
	tx, err := e.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	employee, err := e.EmployeeRepository.GetById(ctx, tx, employeeId)
	if err != nil {
		panic(exception.NewNotFoundException(err.Error()))
	}

	return helpers.ConvertEmployeeResponse(employee)
}

// Update implements EmployeeService.
func (e *EmployeeServiceImpl) Update(ctx context.Context, request web.EmployeeUpdateRequest) web.EmployeeResponse {
	err := e.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx, err := e.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	employee, err := e.EmployeeRepository.GetById(ctx, tx, request.EmployeeId)
	if err != nil {
		panic(exception.NewNotFoundException(err.Error()))
	}

	employee.Fullname = request.Fullname
	employee.Addrees = request.Address

	employee = e.EmployeeRepository.Update(ctx, tx, employee)

	return helpers.ConvertEmployeeResponse(employee)
}
