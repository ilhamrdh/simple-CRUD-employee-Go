package web

type EmployeeCreateRequest struct {
	EmployeeId string `validate:"required,min=4,max=20" json:"employee_id"`
	Fullname   string `validate:"required,min=2,max=50" json:"fullname"`
	Address    string `validate:"required,min=2,max=100" json:"address"`
}

type EmployeeUpdateRequest struct {
	EmployeeId string `validate:"required,min=4,max=20" json:"employee_id"`
	Fullname   string `validate:"required,min=2,max=50" json:"fullname"`
	Address    string `json:"address"`
}
