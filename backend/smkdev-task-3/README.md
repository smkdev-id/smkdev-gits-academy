# **_REST API CRUD Books_**

Books store application built with [Golang](https://go.dev/) using the [Echo framework](github.com/labstack/echo/v4), [SQLite](https://www.sqlite.org/), and [Gorm](https://gorm.io/). The JSON responses are designed in a way that is easy to interpret and use.

Where full URLs are provided in responses they will be rendered as if service
is running on 'http://localhost:8080/'.

## Tech Stack

**Language:** Golang \
**Server:** Echo \
**ORM:** GORM \
**Database:** SQLite

## Project Structure

Here is the structure of the project:

```
smkdev-task-3
├── cmd
│ └── main.go
├── pkg
│ ├── config
│ │ └── app.go
│ ├── controllers
│ │ └── bookController.go
│ ├── database
│ │ ├── bookstore.db
│ │ └── sqlite.go
│ ├── models
│ │ └── book.go
│ ├── routes
│ │ └── booksRoutes.go
│ ├── server
│ │ └── server.go
│ └── utils
│ └── response.go
├── .env
├── go.mod
├── go.sum
└── README.md
```

## Run Locally

Clone the project

```bash
  git clone https://github.com/smkdev-id/smkdev-gits-academy.git
```

Go to the project directory

```bash
  cd smkdev-gits-academy/backend/smkdev-task-3/
```

Start the server

```bash
  go run cmd/main.go (running main.go)
```

URLs are provided in responses they will be rendered as if service
is running on 'http://localhost:8080/'.

## Books Endpoints

These endpoints are related to the Books CRUD operations:

### Get All Book

Path : `/book` \
Method : `GET`

#### Request

```
curl http://localhost:8080/book
```

#### Successful Responses

Code : `200 OK` \
Content examples

```json
{
  "status_code": 200,
  "message": "Successfully get a book",
  "data": [
    {
      "id": "02ad8f1b-3987-45fc-8f93-99ee12991585",
      "title": "My hero Academia",
      "author": "Kohei Horikoshi",
      "publication_year": "2014",
      "price": 21999,
      "stock": 25,
      "isbn": "g7h89jklm34n",
      "synopsis": "Di dunia di mana hampir semua orang memiliki kekuatan super, Izuku Midoriya yang lahir tanpa kekuatan berusaha menjadi pahlawan terbesar.",
      "createdAt": "2024-08-05T00:34:04Z",
      "updatedAt": "2024-08-05T01:28:40.5803878+07:00"
    },
    {
      "id": "de61b9a2-ec49-4999-b975-169c52dce708",
      "title": "Demon Slayer: Kimetsu no Yaiba",
      "author": "Koyoharu Gotouge",
      "publication_year": "2016",
      "price": 24999,
      "stock": 15,
      "isbn": "bz872jhf56hg",
      "synopsis": "Tanjiro Kamado, seorang anak yang menjadi pemburu iblis setelah keluarganya dibunuh oleh iblis dan adiknya berubah menjadi iblis.",
      "createdAt": "2024-08-05T00:50:09.8909977+07:00",
      "updatedAt": "0001-01-01T00:00:00Z"
    }
  ]
}
```

#### Error Responses

Code : `500 Internal Server Error`

### Get A Book

Path : `book/:id` \
Method : `GET`

#### Request

```
curl http://localhost:8080/book/de61b9a2-ec49-4999-b975-169c52dce708
```

#### Successful Responses

Code : `200 OK` \
Content examples

```json
{
  "status_code": 200,
  "message": "Succesfully get a book",
  "data": {
    "id": "de61b9a2-ec49-4999-b975-169c52dce708",
    "title": "Demon Slayer: Kimetsu no Yaiba",
    "author": "Koyoharu Gotouge",
    "publication_year": "2016",
    "price": 24999,
    "stock": 15,
    "isbn": "bz872jhf56hg",
    "synopsis": "Tanjiro Kamado, seorang anak yang menjadi pemburu iblis setelah keluarganya dibunuh oleh iblis dan adiknya berubah menjadi iblis.",
    "createdAt": "2024-08-05T00:50:09.8909977+07:00",
    "updatedAt": "0001-01-01T00:00:00Z"
  }
}
```

#### Error Responses

Code : `500 Internal Server Error` or `404 Not found`

### Create Book

Path : `/book` \
Method : `POST`

#### Request

```
curl http://localhost:8080/book
```

This is a client request to create a book

```json
{
  "title": "Death Note",
  "author": "Tsugumi Ohba",
  "publication_year": "2003",
  "price": 22999,
  "stock": 22,
  "isbn": "r5t7v8n9w1q4",
  "synopsis": "Light Yagami menemukan buku catatan yang memungkinkan dia membunuh siapa saja hanya dengan menulis nama mereka di dalamnya dan berusaha menciptakan dunia tanpa kejahatan."
}
```

#### Successful Responses

Code : `201 Created` \
Content examples

```json
{
  "status_code": 201,
  "message": "Successfully created",
  "data": {
    "id": "d971c83b-9b72-48a9-9959-0a7580a05c91",
    "title": "Death Note",
    "author": "Tsugumi Ohba",
    "publication_year": "2003",
    "price": 22999,
    "stock": 22,
    "isbn": "r5t7v8n9w1q4",
    "synopsis": "Light Yagami menemukan buku catatan yang memungkinkan dia membunuh siapa saja hanya dengan menulis nama mereka di dalamnya dan berusaha menciptakan dunia tanpa kejahatan.",
    "createdAt": "2024-08-05T10:49:38.2672312+07:00",
    "updatedAt": "0001-01-01T00:00:00Z"
  }
}
```

#### Error Responses

Code : `500 Internal Server Error` or `400 Bad request`

### Update A Book

Path : `book/:id` \
Method : `PUT`

#### Request

```
curl http://localhost:8080/book/02ad8f1b-3987-45fc-8f93-99ee12991585
```

This is a client request to update a book

```json
{
  "title": "My hero Academia",
  "author": "Kohei Horikoshi",
  "publication_year": "2014",
  "price": 21999,
  "stock": 25,
  "isbn": "g7h89jklm34n",
  "synopsis": "Di dunia di mana hampir semua orang memiliki kekuatan super, Izuku Midoriya yang lahir tanpa kekuatan berusaha menjadi pahlawan terbesar."
}
```

#### Successful Responses

Code : `200 OK` \
Content examples

```json
{
  "status_code": 200,
  "message": "Successfully updated the book",
  "data": {
    "id": "02ad8f1b-3987-45fc-8f93-99ee12991585",
    "title": "My hero Academia",
    "author": "Kohei Horikoshi",
    "publication_year": "2014",
    "price": 21999,
    "stock": 25,
    "isbn": "g7h89jklm34n",
    "synopsis": "Di dunia di mana hampir semua orang memiliki kekuatan super, Izuku Midoriya yang lahir tanpa kekuatan berusaha menjadi pahlawan terbesar.",
    "createdAt": "2024-08-05T00:34:04Z",
    "updatedAt": "2024-08-05T10:48:28.8751412+07:00"
  }
}
```

#### Error Responses

Code : `500 Internal Server Error` or `400 Bad request`

### Delete A Book

Path : `book/:id` \
Method : `DELETE`

#### Request

```
curl http://localhost:8080/book/d971c83b-9b72-48a9-9959-0a7580a05c91
```

#### Successful Responses

Code : `200 OK` \
Content examples

```json
{
  "status_code": 200,
  "message": "Successfully deleted the book",
  "data": null
}
```

#### Error Response

Code : `500 Internal Server Error`

This README provides a comprehensive guide to understanding, setting up, and interacting with the Books CRUD REST API built with Golang, Echo, GORM, and SQLite3.
