package app

import (
	"database/sql"
	"time"

	"github.com/ilhamrdh/simple-CRUD-employee-Go.git/helpers"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/simple_crud_employee") // use your username and password
	helpers.PanicIfError(err)

	db.SetMaxIdleConns(5)                   // Set max idle connections to 5
	db.SetMaxOpenConns(20)                  // Set max open connections to 20
	db.SetConnMaxLifetime(60 * time.Minute) // Set max lifetime for a connection to 60 minutes
	db.SetConnMaxIdleTime(10 * time.Minute) // Set max idle time for a connection to 10 minutes

	return db
}
