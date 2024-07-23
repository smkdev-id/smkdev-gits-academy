# Todo List API

This is a simple CRUD API for managing a todo list, built using Golang with the Echo framework, GORM for ORM, and SQLite for the database.

## Technologies

- **Golang**: Programming language used for the server-side logic.
- **Echo**: A high-performance web framework for Go.
- **SQLite**: Lightweight SQL database.

### Installation

**Clone the Repository**

```bash
git clone https://github.com/smkdev-id/smkdev-gits-academy.git
```

**Go to Project Repository**

```bash
cd smkdev-gits-academy/backend/simple-crud-todo/
```

**Run the Server**

```bash
go run main.go
```

# API Endpoints

**Fetch Todos /**

Retrieve all todo items.

Response:

```json
[
  {
    "id": 1,
    "name": "name",
    "status": false,
    "created_at": "2024-07-23T05:35:30.8150422+07:00",
    "updated_at": "2024-07-23T05:35:30.8150422+07:00"
  }
]
```

**CreateTodos /**

Create a new task item.

Request Body:

```json
[
  {
    "task": "Saya mau belajar besok",
    "status": false
  }
]
```

Response:

```json
[
  {
    "id": 1,
    "task": "Saya mau belajar besok",
    "status": false,
    "created_at": "2024-07-23T05:35:30.8150422+07:00",
    "updated_at": "2024-07-23T05:35:30.8150422+07:00"
  }
]
```

**UpdateTodos /:id**

Update an existing todo by ID.

Request Body:

```json
[
  {
    "task": "Saya mau belajar besok",
    "status": true
  }
]
```

Response:

```json
[
  {
    "id": 1,
    "task": "Saya mau belajar besok",
    "status": true,
    "created_at": "2024-07-23T05:35:30.8150422+07:00",
    "updated_at": "2024-07-23T05:35:30.8150422+07:00"
  }
]
```

**DeleteTodos /:id**

Delete a specific todo by ID.

Response:

```json
{
  "data": true
}
```