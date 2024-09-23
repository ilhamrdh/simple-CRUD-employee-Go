package routes

import (
	"github.com/ilhamrdh/simple-CRUD-employee-Go.git/controllers"
	"github.com/ilhamrdh/simple-CRUD-employee-Go.git/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRoute(employeeController controllers.EmployeeController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/employees", employeeController.FindAll)
	router.GET("/api/employees/:employeeId", employeeController.FindById)
	router.POST("/api/employees", employeeController.Create)
	router.PUT("/api/employees/:employeeId", employeeController.Update)
	router.DELETE("/api/employees/:employeeId", employeeController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
