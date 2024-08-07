# Bookstore REST API

This project is a Bookstore REST API built with Golang, Echo framework, GORM, and SQLite. It provides endpoints for creating, reading, updating, and deleting books in the bookstore.

## Authors

- [@Tri Krama](https://www.github.com/trikrama)

## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Running the Application](#running-the-application)
- [API Endpoints](#api-endpoints)
- [Project Structure](#project-structure)
- [Configuration](#configuration)
- [Utils](#utils)

## Features

- CRUD operations for books
- Structured logging
- Middleware for CORS, logging, and recovery
- Environment configuration
- Validation of request payloads
- Error handling

## Prerequisites

- [Golang](https://golang.org/dl/) (1.18+)
- [SQLite](https://www.sqlite.org/download.html)
- [Echo](https://echo.labstack.com/)
- [GORM](https://gorm.io/index.html)
- [Git](https://git-scm.com/)
- [Postman](https://www.postman.com/) or any API testing tool

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/smkdev-id/smkdev-gits-academy/tree/tri-krama/Backend/BookStore
    ```
    ```bash
    cd BookStore
    ```

2. Initialize and update Go modules:

    ```bash
    go mod tidy
    ```

3. Create a `.env` file in the root directory and add the following configuration:

    ```env
    PORT=8080
    DB_NAME=bookstore
    ```

## Running the Application

1. Build and run the application:

    ```bash
    go run cmd/main.go
    ```

2. The application will start on port 8080 by default. You can test the API using Postman or any other API testing tool.

## API Endpoints

### Books

- **Get All Books**
  - **URL:** `/api/book`
  - **Method:** `GET`
  - **Success Response:**
    - **Code:** 200
    - **Message:** Request was successful
    - **Content:** 
        ```json 
        [
          {
          "id": 1,
          "title": "Book Title",
          "author": "Author Name",
          "price": 199,
          "created_at": "2024-08-07 11:31:05",
          "updated_at": "2024-08-07 11:31:05"
          },
          {
            "id": 2,
            "title": "Book Title 2",
            "author": "Author Name 2",
            "price": 1999,
            "created_at": "2024-08-07 13:03:47",
            "updated_at": "2024-08-07 13:03:47"
          }
        ]
        ```
- **Get Book by ID**
  - **URL:** `/api/book/:id`
  - **Method:** `GET`
  - **Success Response:**
    - **Code:** 200
    - **Message:** Request was successful
    - **Content:** 
        ```json
        {
        "id": 1,
        "title": "Book Title",
        "author": "Author Name",
        "price": 199,
        "created_at": "2024-08-07 11:31:05",
        "updated_at": "2024-08-07 11:31:05"
        }
        ```
  - **Error Response:**
    - **Code:** 404
    - **Content:** 
        ```json
          {
            "status": "error",
            "message": "Book not found"
          }
        ```

- **Create Book**
  - **URL:** `/api/book`
  - **Method:** `POST`
  - **Header:** `Content-Type: application/json`
  - **Request Body:**
    - **Body:** 
        ```json
          { 
            "title": "Book Title 2", 
            "author": "Author Name 2", 
            "price": 1999
          }
        ```
  - **Success Response:**
    - **Code:** 201
    - **Message:** Request was successful
    - **Content:** 
        ```json
          {
            "status": "success",
            "message": "Request was successful"
          }
        ```
  - **Error Response:**
    - **Code:** 400
    - **Message:** Validation error
    - **Content:** 
        ```json
            {
              "status": "error",
              "message": "Validation error",
              "data": {
                "Author": "Author should be at least 3 characters long",
                "Title": "Title should be at least 3 characters long"
              }
            }
        ```
- **Update Book**
  - **URL:** `/api/book/:id`
  - **Method:** `PUT`
  - **Header:** `Content-Type: application/json`
  - **Request Body:**
    - **Content:** 
        ```json
            { 
              "title": "Book Title 2", 
              "author": "Author Name 2", 
              "price": 100000
            }
        ```
  - **Success Response:**
    - **Code:** 200
    - **Message:** Request was successful
    - **Content:** 
        ```json
            {
              "status": "success",
              "message": "Request was successful"
            }
        ```
  - **Error Response:**
    - **Code:** 404
    - **Content:** 
        ```json
              {
                "status": "error",
                "message": "Book not found"
              }
        ```

- **Delete Book**
  - **URL:** `/api/book/:id`
  - **Method:** `DELETE`
  - **Success Response:**
    - **Code:** 200
    - **Message:** Request was successful
    - **Content:** 
        ```json
            {
              "status": "success",
              "message": "Request was successful"
            }
        ```
  - **Error Response:**
    - **Code:** 404
    - **Content:** 
        ```json
              {
                "status": "error",
                "message": "Book not found"
              }
        ```

## Project Structure
```bash
Bookstore/
├── cmd/
│ └── main.go
├── pkg/
│ ├── config/
│ │ └── app.go
│ ├── controllers/
│ │ └── book-controllers.go
│ ├── models/
│ │ └── book.go
│ ├── routes/
│ │ └── bookstore-routes.go
│ ├── utils/
│ │ ├── config.go
│ │ ├── logger.go
│ │ ├── middleware.go
│ │ ├── response.go
│ │ ├── validation.go
├── .env
├── go.sum
├── go.mod
└── Readme.md
```

## Configuration

The application uses a `.env` file for configuration. Ensure you have the following variables set:

```env
PORT=8080
DB_NAME=bookstore.db
```

## Tech Stack

**Database:** Sqlite

**ORM:** GORM

**Framework:** Echo Golang

## License

*Proyek ini dilisensikan di bawah MIT License.*

Dilarang menjual atau mendistribusikan perangkat lunak ini untuk tujuan komersial tanpa izin tertulis dari kami.

