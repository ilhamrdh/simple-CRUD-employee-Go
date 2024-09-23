package entity

import "time"

type Employees struct {
	EmployeeId string
	Fullname   string
	Addrees    string
	UpdatedAt  time.Time
	CreatedAt  time.Time
}
