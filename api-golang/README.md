# Simple Todos CRUD Application

## Table of Contents
- [Introduction](#introduction)
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Configuration](#configuration)
- [License](#license)
- [Acknowledgements](#acknowledgements)

## Introduction
Welcome to the **Simple Todos CRUD Application**! This project is a basic implementation of a CRUD application for managing todos. It uses Go, the net/http package for handling HTTP requests, GORM as the ORM, and SQLite as the database.

## Features
- Create a new todo
- Retrieve all todos
- Retrieve a single todo by ID
- Update a todo
- Delete a todo

## Prerequisites
- Go 1.16 or higher
- SQLite

## Installation
To install and run this project, follow these steps:

1. Clone the repository:
    ```bash
    git clone https://github.com/your-username/simple-todos-crud.git
    ```

2. Navigate to the project directory:
    ```bash
    cd simple-todos-crud
    ```

3. Install dependencies:
    ```bash
    go mod tidy
    ```

## Usage
To start the application, use the following command:

```bash
go run main.go
```

The application will start on http://localhost:8080.

## API Endpoints
Here are the available API endpoints:

### Create a new todo
URL: 
Method: POST
Request Body:

```json
{
    "title": "Learn Go",
    "description": "Study the basics of Go programming language",
    "is_completed": true
}
```

Response:

```json
{
    "id": 1,
    "title": "Learn Go",
    "description": "Study the basics of Go programming language",
    "is_completed": true
}
```
### Retrieve all todos
URL: 
Method: GET
Response:

```json
[
    {
        "id": 1,
        "title": "Learn Go",
        "description": "Study the basics of Go programming language",
        "is_completed": true
    }
]
```

### Update a todo
URL: /{id}
Method: PUT
Request Body:
```json
{
    "title": "Learn Go and GORM",
    "description": "Study the basics of Go and GORM",
    "is_completed": false
}
```
Reponse:
```json
{
    "id": 1,
    "title": "Learn Go and GORM",
    "description": "Study the basics of Go and GORM",
    "is_completed": false
}
```

### Delete a todo
URL: /{id}
Method: DELETE
Response:
```json

{
    "message": "Todo deleted successfully"
}
```

## License
This project is licensed under the MIT License. See the LICENSE file for details.

## Acknowledgements
We would like to thank the following resources and projects that have helped us in developing this application:

[Go Programming Language](https://go.dev)
[GORM](https://gorm.io/)
[SQLite](https://www.sqlite.org/)
[GORM SQLite](https://pkg.go.dev/gorm.io/driver/sqlite)
