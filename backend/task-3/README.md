# BOOK STORE 📖

Create, Read, Update, Delete (CRUD) application of selling books.

---

## TECH STACK 🧑‍💻

- [Go](https://go.dev/)
- [Echo Framework](https://echo.labstack.com/)
- [Sqlite](https://www.sqlite.org/)
- [Golang Validator](https://github.com/go-playground/validator)
- [Viper](https://github.com/spf13/viper)
- [Testify](https://github.com/stretchr/testify)
- [Mockery](https://github.com/vektra/mockery)
- [Gorm](https://gorm.io/)

---

## FEATURES 🛫

- **FIND ALL BOOKS**: looking all books created have exist in database using `GET` method and paginate.
- **FIND ALL BOOKS SEARCH**: looking all book if exist in database, and search with key & value.
- **FIND BOOK BY ID**: find specific book with id params into URL using `GET` method.
- **FIND BOOK BY ISBN**: find specific boom with ISBN params input to URL using `GET` method.
- **CREATE NEW BOOK**: generate new book for display into store using `POST` method.
- **UPDATE BOOK**: update todo when is_displayed=false to is_completed=true and another using `PUT` method.
- **DELETE TODO BY ID**: delete specific book if you completed with id params into URL using `DELETE` method.
- **DELETE ALL BOOKS**: delete all books if exist in the table using `DELETE` method.

---

- [BOOK STORE 📖](#book-store-)
  - [TECH STACK 🧑‍💻](#tech-stack-)
  - [FEATURES 🛫](#features-)
  - [HOW TO RUN TESTING 🧪](#how-to-run-testing-)
  - [HOW TO RUN APPLICATION 🏃](#how-to-run-application-)
  - [HOW TO RUN USING DOCKER 🏃‍♀️](#how-to-run-using-docker-️)
  - [ROUTER, REQUEST AND RESPONSE TODO 💁](#router-request-and-response-todo-)

---

## HOW TO RUN TESTING 🧪

```bash
go test -v ./tests/...
```

---

## HOW TO RUN APPLICATION 🏃

1. you can setup `env.sh` file, get example from `env.example.sh` file.

```sh
#!bin/bash

export APP_ENV=<your-mode-env-app>
export APP_PORT=<your-port-app>

export DATABASE_PATH=<your-database-path>
```

2. after setup, open your terminal, you can run command :

```bash
source env.sh
```

3. so do this command to run project :

```bash
go run cmd/main.go
```

---

## HOW TO RUN USING DOCKER 🏃‍♀️

1. setup your environment variable in file `env.sh` :

```sh
    #!bin/bash

    export APP_ENV="development"
    export APP_PORT="5000"

    export DATABASE_PATH="bookstore.db"
```

- next, open terminal :

```bash
source env.sh
```

2. go to directory :

```bash
cd build/ci
```

```bash
sudo docker compose up --build
```

- if you want to down/stop all :

```bash
sudo docker compose down -v
```

3. If you success, congratulation 💯

---

## ROUTER, REQUEST AND RESPONSE TODO 💁

1. URL Pattern: `/api/v1/books`

- Find All Books - RESPONSE
  - HTTP methods: `GET`
  - Content-Type: application/json

```json
{
  "status_code": 200,
  "message": "Find all books successfully",
  "data": [
    {
      "id": "925a45f6-754e-434b-a2c6-a5b00c3445a6",
      "isbn": "593-0-50648-231-8",
      "book_title": "title 1",
      "book_description": "desc 1",
      "author": "author 1",
      "year_of_manufacture": 1999,
      "stock": 20,
      "price": 10000,
      "is_displayed": false,
      "created_at": "2024-08-03 05:10:00.641046093 +0700 +0700",
      "updated_at": ""
    },
    {
      "id": "5fc578f3-67d0-4cda-be7a-cf68a240f010",
      "isbn": "700-0-02610-379-9",
      "book_title": "title 1",
      "book_description": "desc 1",
      "author": "author 1",
      "year_of_manufacture": 1999,
      "stock": 20,
      "price": 10000,
      "is_displayed": false,
      "created_at": "2024-08-03 12:27:50.645708528 +0700 +0700",
      "updated_at": ""
    },
    {
      "id": "11985502-8a10-4b69-b2d2-34f43ccb479c",
      "isbn": "798-2-94465-250-0",
      "book_title": "title 1",
      "book_description": "desc 1",
      "author": "author 1",
      "year_of_manufacture": 1999,
      "stock": 20,
      "price": 10000,
      "is_displayed": false,
      "created_at": "2024-08-03 12:33:59.267952551 +0700 +0700",
      "updated_at": ""
    }
  ],
  "pagination": {
    "current_page": 1,
    "page_size": 10,
    "total_data": 3,
    "total_pages": 1
  }
}
```

2. URL Pattern: `/api/v1/books/search`

- Query Params Optional URL:

```json
"isbn": "",
"title": "",
"year": 0,
"author": "",
"price": 0,
"start_date": "", // search when input start date from created the book
"end_date": "", // search when input end date from created the book
"is_displayed": false // search book if book dispalayed in store or not
```

- Find All Books Search - RESPONSE
  - HTTP methods: `GET`
  - Content-Type: application/json

```json
{
  "status_code": 200,
  "message": "Find all books successfully",
  "data": [
    {
      "id": "925a45f6-754e-434b-a2c6-a5b00c3445a6",
      "isbn": "593-0-50648-231-8",
      "book_title": "title 1",
      "book_description": "desc 1",
      "author": "author 1",
      "year_of_manufacture": 1999,
      "stock": 20,
      "price": 10000,
      "is_displayed": false,
      "created_at": "2024-08-03 05:10:00.641046093 +0700 +0700",
      "updated_at": ""
    },
    {
      "id": "5fc578f3-67d0-4cda-be7a-cf68a240f010",
      "isbn": "700-0-02610-379-9",
      "book_title": "title 1",
      "book_description": "desc 1",
      "author": "author 1",
      "year_of_manufacture": 1999,
      "stock": 20,
      "price": 10000,
      "is_displayed": false,
      "created_at": "2024-08-03 12:27:50.645708528 +0700 +0700",
      "updated_at": ""
    }
  ],
  "pagination": {
    "current_page": 1,
    "page_size": 10,
    "total_data": 2,
    "total_pages": 1
  }
}
```

3. URL Pattern: `/api/v1/books`

- Create New Book - REQUEST
  - HTTP methods: `POST`
  - Content-Type: application/json

```json
{
  "title": "title book",
  "description": "description book",
  "author": "author book",
  "year": 2012, // year since the book
  "stock": 20, // stock in store
  "price": 10000 // price the book
}
```

- Create New Book - RESPONSE
  - Content-Type: application/json

```json
{
  "status_code": 201,
  "message": "Create book successfully",
  "data": null,
  "pagination": {
    "current_page": 0,
    "page_size": 0,
    "total_data": 0,
    "total_pages": 0
  }
}
```

4. URL Pattern: `/api/v1/books/:id`

- Find Book By Id - RESPONSE
  - HTTP methods: `GET`
  - Content-Type: application/json

```json
{
  "status_code": 200,
  "message": "Find book by id successfully",
  "data": {
    "id": "925a45f6-754e-434b-a2c6-a5b00c3445a6",
    "isbn": "593-0-50648-231-8",
    "book_title": "title 1",
    "book_description": "desc 1",
    "author": "author 1",
    "year_of_manufacture": 1999,
    "stock": 20,
    "price": 10000,
    "is_displayed": false,
    "created_at": "2024-08-03 05:10:00.641046093 +0700 +0700",
    "updated_at": ""
  },
  "pagination": {
    "current_page": 0,
    "page_size": 0,
    "total_data": 0,
    "total_pages": 0
  }
}
```

5. URL Pattern: `/api/v1/books/isbn/:isbn`

- Find Book By ISBN - RESPONSE
  - HTTP methods: `GET`
  - Content-Type: application/json

```json
{
  "status_code": 200,
  "message": "Find book by id successfully",
  "data": {
    "id": "925a45f6-754e-434b-a2c6-a5b00c3445a6",
    "isbn": "593-0-50648-231-8",
    "book_title": "title 1",
    "book_description": "desc 1",
    "author": "author 1",
    "year_of_manufacture": 1999,
    "stock": 20,
    "price": 10000,
    "is_displayed": false,
    "created_at": "2024-08-03 05:10:00.641046093 +0700 +0700",
    "updated_at": ""
  },
  "pagination": {
    "current_page": 0,
    "page_size": 0,
    "total_data": 0,
    "total_pages": 0
  }
}
```

6. URL Pattern: `/api/v1/books/:id`

- Update Book - REQUEST
  - HTTP methods: `PUT`
  - Content-Type: application/json

```json
{
  "title": "title 2",
  "description": "desc 2",
  "author": "author 2",
  "year": 2000,
  "stock": 55,
  "is_displayed": true,
  "price": 105000
}
```

- Update Book - RESPONSE
  - Content-Type: application/json

```json
{
  "status_code": 200,
  "message": "Update book successfully",
  "data": null,
  "pagination": {
    "current_page": 0,
    "page_size": 0,
    "total_data": 0,
    "total_pages": 0
  }
}
```

7. URL Pattern: `/api/v1/books/:id`

- DELETE Book By ID - RESPONSE
  - HTTP methods: `DELETE`
  - Content-Type: application/json

```json
{
  "status_code": 200,
  "message": "Delete book successfully",
  "data": null,
  "pagination": {
    "current_page": 0,
    "page_size": 0,
    "total_data": 0,
    "total_pages": 0
  }
}
```

8. URL Pattern: `/api/v1/books`

- DELETE All Books - RESPONSE
  - HTTP methods: `DELETE`
  - Content-Type: application/json

```json
{
  "status_code": 200,
  "message": "Delete all books successfully",
  "data": null,
  "pagination": {
    "current_page": 0,
    "page_size": 0,
    "total_data": 0,
    "total_pages": 0
  }
}
```

---

| CONGRATULATION 👍 |
| ----------------- |
