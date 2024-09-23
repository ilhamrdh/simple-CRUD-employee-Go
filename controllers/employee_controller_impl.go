package controllers

import (
	"net/http"

	"github.com/ilhamrdh/simple-CRUD-employee-Go.git/helpers"
	"github.com/ilhamrdh/simple-CRUD-employee-Go.git/models/web"
	"github.com/ilhamrdh/simple-CRUD-employee-Go.git/services"
	"github.com/julienschmidt/httprouter"
)

type EmployeeControllerImpl struct {
	EmployeeService services.EmployeeService
}

func NewEmployeeController(employeeService services.EmployeeService) EmployeeController {
	return &EmployeeControllerImpl{
		EmployeeService: employeeService,
	}
}

// Create implements EmployeeController.
func (e *EmployeeControllerImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	employeeRequest := web.EmployeeCreateRequest{}
	helpers.ReadFormRequestBody(r, &employeeRequest)

	employeeResponse := e.EmployeeService.Create(r.Context(), employeeRequest)
	response := web.WebResponse{
		Code:   201,
		Status: "OK",
		Data:   employeeResponse,
	}

	helpers.WriteJsonResponse(w, response)
}

// Delete implements EmployeeController.
func (e *EmployeeControllerImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	employeeId := p.ByName("employeeId")

	e.EmployeeService.Delete(r.Context(), employeeId)
	helpers.WriteJsonResponse(w, web.WebResponse{
		Code:   200,
		Status: "OK",
	})
}

// FindAll implements EmployeeController.
func (e *EmployeeControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	employees := e.EmployeeService.GetAll(r.Context())

	response := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   employees,
	}

	helpers.WriteJsonResponse(w, response)
}

// FindById implements EmployeeController.
func (e *EmployeeControllerImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	employeeId := p.ByName("employeeId")

	employee := e.EmployeeService.GetById(r.Context(), employeeId)
	response := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   employee,
	}

	helpers.WriteJsonResponse(w, response)
}

// Update implements EmployeeController.
func (e *EmployeeControllerImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	employeeId := p.ByName("employeeId")
	employeeRequest := web.EmployeeUpdateRequest{}
	helpers.ReadFormRequestBody(r, &employeeRequest)

	employeeRequest.EmployeeId = employeeId

	employeeResponse := e.EmployeeService.Update(r.Context(), employeeRequest)
	response := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   employeeResponse,
	}

	helpers.WriteJsonResponse(w, response)
}
