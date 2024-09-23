package main

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql" // init mysql driver
	"github.com/ilhamrdh/simple-CRUD-employee-Go.git/app"
	"github.com/ilhamrdh/simple-CRUD-employee-Go.git/controllers"
	"github.com/ilhamrdh/simple-CRUD-employee-Go.git/helpers"
	"github.com/ilhamrdh/simple-CRUD-employee-Go.git/repositories"
	"github.com/ilhamrdh/simple-CRUD-employee-Go.git/routes"
	"github.com/ilhamrdh/simple-CRUD-employee-Go.git/services"
)

func main() {

	// init database
	db := app.NewDB()

	// init validator
	validate := validator.New()

	// init repository, service and controller
	employeeRepository := repositories.NewEmployeeRepository()
	employeeService := services.NewEmployeeService(employeeRepository, db, validate)
	employeeController := controllers.NewEmployeeController(employeeService)

	// init router
	router := routes.NewRoute(employeeController)

	// init server
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	fmt.Println("Server started at localhost:8080")

	err := server.ListenAndServe()
	helpers.PanicIfError(err)

}
