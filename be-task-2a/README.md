Todo API Project
# Todo API

This project is a simple Todo API built with Go, GORM, and PostgreSQL. It provides endpoints for creating, reading, updating, and deleting todo items.

## Prerequisites
- Go 1.16 or later
- PostgreSQL
- Visual Studio Code (VS Code)
- Thunder Client extension for VS Code

## Setup
1. Clone the Repository
    ```bash
    git clone https://github.com/HerlyRyan/be-task-2a.git
    cd be-task-2a
    ```

2. Configure the Database
    - Ensure you have PostgreSQL installed and running.
    - Update the `config.go` file with your PostgreSQL database credentials.

    ```go
    // config.go
    package db

    import (
         "fmt"
         "log"
         "gorm.io/driver/postgres"
         "gorm.io/gorm"
    )

    var (
         UNAMEDB string = "postgres" // Use the correct PostgreSQL user
         PASSDB  string = "your_password" // Use the correct password
         HOSTDB  string = "localhost"
         DBNAME  string = "todo-db"
    )

    func Connectdb() *gorm.DB {
         connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", HOSTDB, UNAMEDB, PASSDB, DBNAME)
         db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
         if err != nil {
              log.Fatal(err)
         }
         return db
    }
    ```

3. Create the Database and Tables
    - Connect to your PostgreSQL instance and create the database:

    ```sql
    CREATE DATABASE "todo-db";
    ```

    - Run the application to automatically migrate the todo table schema:

    ```bash
    go run main.go
    ```

4. Running the Application
    ```bash
    go run main.go
    ```
    The application will start on http://localhost:8080.

## API Endpoints
- `GET /`
  ```sh
  curl -X GET http://localhost:8080/
  ```

- `POST /todos`
  ```sh
  curl -X POST http://localhost:8080/todos \
  -H "Content-Type: application/json" \
  -d '{
          "task": "New Task",
          "due_date": "2024-07-25T00:00:00Z",
          "status": "pending",
          "priority": "medium"
        }'
  ```

- `GET /todos`
  ```sh
  curl -X GET http://localhost:8080/todos
  ```

- `GET /todos/`
  ```sh
  curl -X GET http://localhost:8080/todos/1
  ```

- `PUT /todos/`
  ```sh
  curl -X PUT http://localhost:8080/todos/1 \
  -H "Content-Type: application/json" \
  -d '{
          "task": "Updated Task",
          "due_date": "2024-07-25T00:00:00Z",
          "status": "in_progress",
          "priority": "high"
        }'
  ```

- `DELETE /todos/`
  ```sh
  curl -X DELETE http://localhost:8080/todos/1
  ```

## Testing with Thunder Client
- Install Thunder Client Extension:
  1. Open VS Code.
  2. Go to the Extensions view by clicking on the Extensions icon in the Activity Bar on the side of the window or pressing Ctrl+Shift+X.
  3. Search for "Thunder Client" and click "Install".

- Create a New Collection:
  1. Click on the Thunder Client icon in the Activity Bar.
  2. Click the "New Collection" button.
  3. Name your collection, e.g., "Todo API Tests".

- Create Requests:
  - GET /
     1. Click "New Request".
     2. Set the method to GET.
     3. Set the URL to http://localhost:8080/.
     4. Click "Send".

  - POST /todos
     1. Click "New Request".
     2. Set the method to POST.
     3. Set the URL to http://localhost:8080/todos.
     4. Click on the "Body" tab.
     5. Select JSON as the body type.
     6. Enter the following JSON payload:
         ```json
         {
            "task": "New Task",
            "due_date": "2024-07-25T00:00:00Z",
            "status": "pending",
            "priority": "medium"
         }
         ```
     7. Click "Send".

  - GET /todos
     1. Click "New Request".
     2. Set the method to GET.
     3. Set the URL to http://localhost:8080/todos.
     4. Click "Send".

  - GET /todos/
     1. Click "New Request".
     2. Set the method to GET.
     3. Set the URL to http://localhost:8080/todos/1.
     4. Click "Send".

  - PUT /todos/
     1. Click "New Request".
     2. Set the method to PUT.
     3. Set the URL to http://localhost:8080/todos/1.
     4. Click on the "Body" tab.
     5. Select JSON as the body type.
     6. Enter the following JSON payload:
         ```json
         {
            "task": "Updated Task",
            "due_date": "2024-07-25T00:00:00Z",
            "status": "in_progress",
            "priority": "high"
         }
         ```
     7. Click "Send".

  - DELETE /todos/
     1. Click "New Request".
     2. Set the method to DELETE.
     3. Set the URL to http://localhost:8080/todos/1.
     4. Click "Send".

- Save Requests:
  - After configuring each request, save it in your "Todo API Tests" collection by clicking the save icon or using the appropriate option in the Thunder Client interface.

- Run All Requests:
  - Open the collection and click on the "Run All" button to execute all the requests sequentially.

## License
This project is licensed under the MIT License.