# TODO APP 📑

Create, Read, Update, Delete (CRUD) application of todo.

---

## TECH STACK 🧑‍💻

- [Go](https://go.dev/)
- [Golang Migrate](https://github.com/golang-migrate/migrate)
- [MySQL](https://github.com/go-sql-driver/mysql)
- [Golang Validator](https://github.com/go-playground/validator)
- [Viper](https://github.com/spf13/viper)
- [Testify](https://github.com/stretchr/testify)
- [Mockery](https://github.com/vektra/mockery)
- [Go SQL Mock](https://github.com/DATA-DOG/go-sqlmock)

---

## FEATURES 🛫

- **FIND ALL TODOS**: looking all todos created have exist in database using `GET` method.
- **FIND BY ID**: find specific todo with id params into URL using `GET` method.
- **CREATE TODO**: generate new todo with input title and description using `POST` method.
- **UPDATE TODO**: update todo when is_completed=false to is_completed=true, title and description using `PUT` method.
- **DELETE TODO**: delete specific todo if you completed with id params into URL using `DELETE` method.

---

- [TODO APP 📑](#todo-app-)
  - [TECH STACK 🧑‍💻](#tech-stack-)
  - [FEATURES 🛫](#features-)
  - [HOW TO RUN TESTING 🧪](#how-to-run-testing-)
  - [HOW TO RUN APPLICATION 🏃](#how-to-run-application-)
  - [HOW TO RUN USING DOCKER 🏃‍♀️](#how-to-run-using-docker-️)
  - [ROUTER, REQUEST AND RESPONSE TODO 💁](#router-request-and-response-todo-)
  - [ERROR RESPONSE TODO :ambulance:](#error-response-todo-ambulance)

---

## HOW TO RUN TESTING 🧪

```bash
go test -v ./test/...
```

---

## HOW TO RUN APPLICATION 🏃

1. you can setup `env.sh` file, get example from `env.example.sh` file.

```sh
#!bin/bash

# DATABASE
export DB_DRIVER=mysql
export DB_HOST=localhost
export DB_PORT=3306
export DB_USER=<your-db-user>
export DB_PASSWORD=<your-db-password>
export DB_NAME=<your-db-name>

# APPLICATION
export APP_PORT=":8080"
```

2. after setup, open your terminal, you can run command :

```bash
source env.sh
```

3. generate migrate create ⤵️ :

```bash
make migrate-crete
```

4. setup your table in `db/migration/<date-exec>_task2.up.sql`.

```sql
CREATE TABLE todo (
    id VARCHAR(55) PRIMARY KEY NOT NULL,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(555) NOT NULL,
    is_completed BOOLEAN NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
) ENGINE=INNODB;
```

5. if you have written table, you can exec with command :

```bash
make migrate-up
```

- then, if you want down/drop your table generated :

```bash
make migrate-down
```

6. so do this command :

```bash
go run main.go
```

---

## HOW TO RUN USING DOCKER 🏃‍♀️

1. setup your environment variable in file `env.sh` :

```sh
#!bin/bash

# DB CONTAINER
export DB_DRIVER=mysql
export DB_HOST=localhost
export DB_PORT=3306
export DB_USER=root
export DB_PASSWORD=123
export DB_NAME=task2

# APPLICATION
export APP_PORT=":8080"
```

- next, open terminal :

```bash
source env.sh
```

2. run with command :

```bash
sudo docker compose build
```

```bash
sudo docker compose up -d
```

- if you want to down/stop all :

```bash
sudo docker compose down -v
```

3. you can inside to container `mysql-container`, using command :

```bash
sudo docker exec -it mysql-container /bin/bash
```

- if you inside `mysql-container`, then go to `mysql CLI`, exec command :

```bash
home# mysql -u root -p
home# password:123
```

4. inside to `mysql`, so check your table exist :

```bash
mysql>  SHOW DATABASES;
        USE task2;
        SHOW TABLES;
```

5. check your table name is `todo` must _exist_ in database 💯

---

## ROUTER, REQUEST AND RESPONSE TODO 💁

1. URL Pattern: `/api/v1/todos`

- Find All Todos - RESPONSE
  - HTTP methods: `GET`
  - Content-Type: application/json

```json
{
  "status_code": 200,
  "message": "Find all todos successfully",
  "data": [
    {
      "id": "c216239d-ad00-46f9-a9c5-632ec197befb",
      "title": "title",
      "description": "description",
      "is_completed": false,
      "created_at": "2024-07-20 05:57:38 +0700 WIB",
      "updated_at": ""
    },
    {
      "id": "c7a2e5f9-bd36-4156-bc9c-e7ea9417f096",
      "title": "title",
      "description": "description",
      "is_completed": true,
      "created_at": "2024-07-20 05:57:38 +0700 WIB",
      "updated_at": "2024-07-20 05:36:46 +0700 WIB"
    },
    {
      "id": "c944cf53-2468-48a4-9231-7367bbd2c75f",
      "title": "title",
      "description": "description",
      "is_completed": false,
      "created_at": "2024-07-20 05:36:46 +0700 WIB",
      "updated_at": ""
    }
  ]
}
```

2. URL Pattern: `/api/v1/todos`

- Create Todo - REQUEST
  - HTTP methods: `POST`
  - Content-Type: application/json

```json
{
  "title": "title",
  "description": "description"
}
```

- Create Todo - RESPONSE
  - Content-Type: application/json

```json
{
  "status_code": 200,
  "message": "Create todo successfully",
  "data": {
    "id": "c7a2e5f9-bd36-4156-bc9c-e7ea9417f096",
    "title": "title",
    "description": "description",
    "is_completed": false,
    "created_at": "2024-07-19 22:22:36.852607593 +0700 WIB",
    "updated_at": ""
  }
}
```

3. URL Pattern: `/api/v1/todos/:id`

- Find Todo By Id - RESPONSE
  - HTTP methods: `GET`
  - Content-Type: application/json

```json
{
  "status_code": 200,
  "message": "Find todo successfully",
  "data": {
    "id": "c7a2e5f9-bd36-4156-bc9c-e7ea9417f096",
    "title": "title",
    "description": "description",
    "is_completed": false,
    "created_at": "2024-07-19 15:22:37 +0700 WIB",
    "updated_at": ""
  }
}
```

4. URL Pattern: `/api/v1/todos/:id`

- Update Todo - REQUEST
  - HTTP methods: `PUT`
  - Content-Type: application/json

```json
{
  "title": "title",
  "description": "description",
  "is_completed": true
}
```

- Update Todo - RESPONSE
  - Content-Type: application/json

```json
{
  "status_code": 200,
  "message": "Update todo successfully",
  "data": {
    "id": "c7a2e5f9-bd36-4156-bc9c-e7ea9417f096",
    "title": "title",
    "description": "description",
    "is_completed": true,
    "created_at": "2024-07-19 15:22:37 +0700 WIB",
    "updated_at": "2024-07-19 22:35:20.785811372 +0700 WIB"
  }
}
```

5. URL Pattern: `/api/v1/todos/:id`

- DELETE Todo - RESPONSE
  - HTTP methods: `DELETE`
  - Content-Type: application/json

```json
{
  "status_code": 200,
  "message": "Delete todo successfully",
  "data": null
}
```

---

## ERROR RESPONSE TODO :ambulance:

1. _CREATE TODO_

- error data type, if your input inside to body not correct.

```json
{
  "status_code":400,
  "message":"Failed to create todo",
  "error":"Invalid request body"
}
json: cannot unmarshal number into Go struct field TodoCreateRequest.title of type string
```

- error bad request or validation, if you inside to body not correct.

```json
{
  "status_code":400,
  "message":"Failed to create todo",
  "error":"Validation error"
}
Key: 'TodoCreateRequest.Title' Error:Field validation for 'Title' failed on the 'min' tag

```

- other error

```json
Internal Server Error
```

2. _FIND ALL TODOS_

- error when find all todos.

```json
Internal Server Error
```

3. _FIND TODO BY ID_

- error not found todo when id not correct.

```json
{"status_code":404,"message":"Failed to find todo","error":"Todo not found"}
sql: no rows in result set
```

- other error.

```json
Internal Server Error
```

4. _UPDATE TODO_

- error not found todo when id not correct.

```json
{"status_code":404,"message":"Failed to find todo","error":"Todo not found"}
sql: no rows in result set
```

- error data type, if your input inside to body not correct.

```json
{
  "status_code":400,
  "message":"Failed to update todo",
  "error":"Invalid request body"
}
json: cannot unmarshal number into Go struct field TodoUpdateRequest.title of type string
```

- error bad request or validation, if you inside to body not correct.

```json
{
  "status_code":400,
  "message":"Failed to update todo",
  "error":"Validation error"
}
Key: 'TodoUpdateRequest.Title' Error:Field validation for 'Title' failed on the 'min' tag

```

- other error

```json
Internal Server Error
```

5. _DELETE TODO_

- error not found todo when id not correct.

```json
{"status_code":404,"message":"Failed to find todo","error":"Todo not found"}
sql: no rows in result set
```

- other error

```json
Internal Server Error
```

| CONGRATULATION 👍 |
| ----------------- |
