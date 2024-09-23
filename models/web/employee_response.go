package web

type EmployeeResponse struct {
	EmployeeId string `json:"employee_id"`
	Fullname   string `json:"fullname"`
	Address    string `json:"address"`
}
