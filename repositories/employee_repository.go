package repositories

import (
	"context"
	"database/sql"

	"github.com/ilhamrdh/simple-CRUD-employee-Go.git/models/entity"
)

type EmployeeRepository interface {
	Save(ctx context.Context, tx *sql.Tx, employee entity.Employees) entity.Employees
	Update(ctx context.Context, tx *sql.Tx, employee entity.Employees) entity.Employees
	Delete(ctx context.Context, tx *sql.Tx, employee entity.Employees)
	GetAll(ctx context.Context, tx *sql.Tx) []entity.Employees
	GetById(ctx context.Context, tx *sql.Tx, employeeId string) (entity.Employees, error)
}
