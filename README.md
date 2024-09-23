# Employee Management API

The Employee Management API is a Go (Golang)-based application that allows users to manage employee data through several REST API endpoints. This application supports basic CRUD (Create, Read, Update, Delete) operations on employee data.

## Key Features

-   Add new employee data
-   View all employees
-   Search for employees by ID
-   Update employee data
-   Delete employee data

## Prerequisites

Before starting, ensure you have the following installed:

-   [Golang](https://golang.org/doc/install) (latest version)
-   MySQL or a compatible database system

## Installation Steps

### 1. Clone the Repository

Clone this repository to your local machine:

```bash
git clone https://github.com/ilhamrdh/simple-CRUD-employee-Go.git
cd simple-CRUD-employee-Go
```

### 2. Import the Database

Import the provided database SQL file located in the database/simple_crud_employee.sql

### 3. Install Dependencies

```bash
go mod tidy
```

### 4 Run the Application

To run the application, use the following command:

```bash
go run main.go
```

## API Endpoints

Below is a list of the available API endpoints:

### 1. **GET /api/employees**

-   **Description**: Fetch a list of all employees.
-   **Response**: An array of employee objects in JSON format.

**Example Response**:

```json
{
    "code": 200,
    "status": "OK",
    "data": [
        {
            "employee_id": "1001",
            "fullname": "Budi",
            "address": "Jakarta"
        },
        {
            "employee_id": "1002",
            "fullname": "Adi",
            "address": "Jakarta"
        },
        {
            "employee_id": "1003",
            "fullname": "Muhammad",
            "address": "Tangerang"
        }
    ]
}
```

### 2. **GET /api/employees/{employeeId}**

-   **Description**: Fetch an employee by their ID.
-   **Parameter**: `employeeId` - The ID of the employee.
-   **Response**: A single employee object in JSON format.

**Example Response**:

```json
{
    "code": 200,
    "status": "OK",
    "data": {
        "employee_id": "1001",
        "fullname": "Budi",
        "address": "Jakarta"
    }
}
```

### 3. **POST /api/employees**

-   **Description**: Create a new employee record.
-   **Request Body (JSON)**:

```json
{
    "employee_id": "1001",
    "fullname": "Budi",
    "address": "Jakarta"
}
```

-   **Response**: The newly created employee object.

**Example Response**:

```json
{
    "code": 201,
    "status": "OK",
    "data": {
        "employee_id": "1001",
        "fullname": "Budi",
        "address": "Jakarta"
    }
}
```

### 4. **PUT /api/employees/{employeeId}**

-   **Description**: Update an existing employee's data.
-   **Parameter**: `employeeId` - The ID of the employee to update.
-   **Request Body (JSON)**:

```json
{
    "fullname": "Budi",
    "address": "Bandung"
}
```

-   **Response**: The updated employee object.

**Example Response**:

```json
{
    "code": 200,
    "status": "OK",
    "data": {
        "employee_id": "1001",
        "fullname": "Budi",
        "address": "Bandung"
    }
}
```

### 5. **DELETE /api/employees/{employeeId}**

-   **Description**: Delete an employee by their ID.
-   **Parameter**: `employeeId` - The ID of the employee to delete.
-   **Response**: A confirmation message.

**Example Response**:

```json
{
    "code": 200,
    "status": "OK",
    "data": null
}
```

## Dependencies

-   `github.com/go-sql-driver/mysql`: MySQL driver for Go.
-   `github.com/julienschmidt/httprouter`: High-performance HTTP router.
-   `github.com/go-playground/validator/v10`: Data validation package.
