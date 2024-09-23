package repositories

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/ilhamrdh/simple-CRUD-employee-Go.git/helpers"
	"github.com/ilhamrdh/simple-CRUD-employee-Go.git/models/entity"
)

type EmployeeRepositoryImpl struct{}

func NewEmployeeRepository() EmployeeRepository {
	return &EmployeeRepositoryImpl{}
}

// Delete implements EmployeeRepository.
func (e *EmployeeRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, employee entity.Employees) {
	SQL := "DELETE FROM employees WHERE employee_id = ?"
	_, err := tx.ExecContext(ctx, SQL, employee.EmployeeId)
	helpers.PanicIfError(err)
}

// GetAll implements EmployeeRepository.
func (e *EmployeeRepositoryImpl) GetAll(ctx context.Context, tx *sql.Tx) []entity.Employees {
	SQL := "SELECT employee_id, fullname, address FROM employees"
	rows, err := tx.QueryContext(ctx, SQL)
	helpers.PanicIfError(err)
	defer rows.Close()

	var employees []entity.Employees
	for rows.Next() {
		employee := entity.Employees{}
		err := rows.Scan(&employee.EmployeeId, &employee.Fullname, &employee.Addrees)
		helpers.PanicIfError(err)
		employees = append(employees, employee)
	}
	return employees
}

// GetById implements EmployeeRepository.
func (e *EmployeeRepositoryImpl) GetById(ctx context.Context, tx *sql.Tx, employeeId string) (entity.Employees, error) {
	SQL := "SELECT employee_id, fullname, address FROM employees WHERE employee_id = ?"
	rows, err := tx.QueryContext(ctx, SQL, employeeId)
	helpers.PanicIfError(err)
	defer rows.Close()

	var employee entity.Employees
	if rows.Next() {
		err := rows.Scan(&employee.EmployeeId, &employee.Fullname, &employee.Addrees)
		helpers.PanicIfError(err)
		return employee, nil
	} else {
		return employee, errors.New("employee not found")
	}
}

// Save implements EmployeeRepository.
func (e *EmployeeRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, employee entity.Employees) entity.Employees {
	SQL := "INSERT INTO employees (employee_id, fullname, address) VALUES (?, ?, ?)"
	_, err := tx.ExecContext(ctx, SQL, employee.EmployeeId, employee.Fullname, employee.Addrees)
	helpers.PanicIfError(err)

	return employee
}

// Update implements EmployeeRepository.
func (e *EmployeeRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, employee entity.Employees) entity.Employees {
	SQL := "UPDATE employees SET fullname = ?, address = ? WHERE employee_id = ?"
	_, err := tx.ExecContext(ctx, SQL, employee.Fullname, employee.Addrees, employee.EmployeeId)
	helpers.PanicIfError(err)

	employee.UpdatedAt = time.Now()

	return employee
}
