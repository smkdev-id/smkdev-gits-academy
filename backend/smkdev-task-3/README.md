# Bookstore API

The Bookstore API is a RESTful API for managing book sales, built using Golang, Echo, GORM, and SQLite.

## Technologies

- **Golang**: The primary programming language for the backend.
- **Echo**: A web framework used to build the HTTP server.
- **GORM**: ORM (Object-Relational Mapping) used to manage interactions with the database.
- **SQLite**: The database used to store book data.
- **Go Modules**: Dependency management system for Go.

## Features

- Get all books
- Create book
- Get detail book by id
- Update book information
- Delete book

## Project Structure
```text
smkdev-task-3/
├── cmd/
│ └── main.go
├── pkg/
│ ├── config/
│ │ └── app.go
│ │ └── bookstore.db
│ ├── controllers/
│ │ └── bookcontrollers.go
│ ├── models/
│ │ └── book.go
│ ├── routes/
│ │ └── bookstores.go
│ └── utils/
│ └── utils.go
├── go.mod
└── go.sum
```

## Installation

Clone the repository

```sh
https://github.com/smkdev-id/smkdev-gits-academy/tree/muhamad-randy-andrian/backend/smkdev-task-3

cd smkdev-task-3
```
Install dependencies
```sh
go mod tidy
```

## Running the Application

Run the following command to start the server
```sh
go run cmd/main.go
```
The server will run at http://localhost:8080

## Endpoints

**Get Books**

- URL: /book
- Method: GET
- Response:

```json
{
    "meta": {
        "message": "Get Books",
        "code": 200,
        "status": "success"
    },
    "data": [
        {
            "id": 1,
            "title": "Is It Bad Or Good Habits Motivasi",
            "description": "Buku Is It Bad Or Good Habbits membahas tentang proses berfikir manusia dan mengenali diri untuk melangkah dari kebiasaan kecil menjadi kebiasaan mengantarkan Anda menuju kesuksesan.",
            "page_count": 134,
            "price": "Rp 30.550",
            "author": "Sabrina Ara",
            "quantity": 1306,
            "created_at": "07 Januari 2024 13:07",
            "updated_at": "07 Januari 2024 13:07"
        },
        {
            "id": 2,
            "title": "Atomic Habits",
            "description": "Orang mengira ketika Anda ingin mengubah hidup, Anda perlu memikirkan hal-hal besar. Namun pakar kebiasaan terkenal kelas dunia James Clear telah menemukan sebuah cara lain. Ia tahu bahwa perubahan nyata berasal dari efek gabungan ratusan keputusan kecil—dari mengerjakan dua push-up sehari, bangun lima menit lebih awal, sampai menahan sebentar hasrat untuk menelepon. Ia menyebut semua tadi atomic habits.",
            "page_count": 352,
            "price": "Rp 91.800",
            "author": "James Clear",
            "quantity": 21,
            "created_at": "07 Januari 2024 13:12",
            "updated_at": "07 Januari 2024 13:12"
        }
    ]
}
```
- Response Error:
  - Book Not Found

    ```json
    {
    "meta": {
        "message": "Data Not Found",
        "code": 404,
        "status": "error"
    },
    "data": null
    }
    ```

**Create Book**
- URL: /book
- Method: POST
- Request Body:
```json
{
    "title": "Is It Bad Or Good Habits Motivasi",
    "description": "Buku Is It Bad Or Good Habbits membahas tentang proses berfikir manusia dan mengenali diri untuk melangkah dari kebiasaan kecil menjadi kebiasaan mengantarkan Anda menuju kesuksesan.",
    "page_count": 134,
    "Author": "Sabrina Ara",
    "Price": 30550,
    "Quantity": 1306
}
```
- Response

```json
{
    "meta": {
        "message": "Success Create Book",
        "code": 200,
        "status": "success"
    },
     {
            "id": 1,
            "title": "Is It Bad Or Good Habits Motivasi",
            "description": "Buku Is It Bad Or Good Habbits membahas tentang proses berfikir manusia dan mengenali diri untuk melangkah dari kebiasaan kecil menjadi kebiasaan mengantarkan Anda menuju kesuksesan.",
            "page_count": 134,
            "price": "Rp 30.550",
            "author": "Sabrina Ara",
            "quantity": 1306,
            "created_at": "07 Januari 2024 13:07",
            "updated_at": "07 Januari 2024 13:07"
        }
}
```
- Response Error:
   - Validation Failed:
   ```json
    {
    "meta": {
        "message": "Key: 'Book.Title' Error:Field validation for 'Title' failed on the 'required' tag\nKey: 'Book.Description' Error:Field validation for 'Description' failed on the 'required' tag\nKey: 'Book.PageCount' Error:Field validation for 'PageCount' failed on the 'required' tag\nKey: 'Book.Author' Error:Field validation for 'Author' failed on the 'required' tag\nKey: 'Book.Price' Error:Field validation for 'Price' failed on the 'required' tag\nKey: 'Book.Quantity' Error:Field validation for 'Quantity' failed on the 'required' tag",
        "code": 400,
        "status": "error"
    },
    "data": null
    }
   ```

**Get Book by Id**

- URL: /book/:id
- Method: GET
- Response:

```json
{
    "meta": {
        "message": "Get Book by Id",
        "code": 200,
        "status": "success"
    },
    "data": {
        "id": 1,
        "title": "Is It Bad Or Good Habits Motivasi",
        "description": "Buku Is It Bad Or Good Habbits membahas tentang proses berfikir manusia dan mengenali diri untuk melangkah dari kebiasaan kecil menjadi kebiasaan mengantarkan Anda menuju kesuksesan.",
        "page_count": 134,
        "price": "Rp 30.550",
        "author": "Sabrina Ara",
        "quantity": 1306,
        "created_at": "07 Januari 2024 13:07",
        "updated_at": "07 Januari 2024 13:07"
    }
}
```
- Response Error
  - Book by Id Not Found
  ``` json
    {
    "meta": {
        "message": "Data Not Found",
        "code": 404,
        "status": "error"
    },
    "data": null
    }
  ```

***Update Book***

- URL: /book/:id
- Method: PUT
- Request Body:

```json
{
    "title": "Is It Good Or Bad Habits Motivasi",
    "description": "Buku Is It Good Or Bad Habbits membahas tentang proses berfikir manusia dan mengenali diri untuk melangkah dari kebiasaan kecil menjadi kebiasaan mengantarkan Anda menuju kesuksesan.",
    "page_count": 134,
    "Author": "Ara Sabrina",
    "Price": 30550,
    "Quantity": 1306
}
```

- Response

```json
{
    "meta": {
        "message": "Success Update Book",
        "code": 200,
        "status": "success"
    },
    "data": {
        "id": 1,
        "title": "Is It Good Or Bad Habits Motivasi",
        "description": "Buku Is It Good Or Bad Habbits membahas tentang proses berfikir manusia dan mengenali diri untuk melangkah dari kebiasaan kecil menjadi kebiasaan mengantarkan Anda menuju kesuksesan.",
        "page_count": 134,
        "price": "Rp 30.550",
        "author": "Ara Sabrina",
        "quantity": 1306,
        "created_at": "07 Januari 2024 13:07",
        "updated_at": "07 Januari 2024 13:42"
    }
}
```
- Response Error
  - Book by Id Not Found
  ```json
    {
    "meta": {
        "message": "Data Not Found",
        "code": 404,
        "status": "error"
    },
    "data": null
    }
  ```
  - Validation Failed

  ```json
    {
    "meta": {
        "message": "Key: 'Book.Title' Error:Field validation for 'Title' failed on the 'required' tag\nKey: 'Book.Description' Error:Field validation for 'Description' failed on the 'required' tag\nKey: 'Book.PageCount' Error:Field validation for 'PageCount' failed on the 'required' tag\nKey: 'Book.Author' Error:Field validation for 'Author' failed on the 'required' tag\nKey: 'Book.Price' Error:Field validation for 'Price' failed on the 'required' tag\nKey: 'Book.Quantity' Error:Field validation for 'Quantity' failed on the 'required' tag",
        "code": 400,
        "status": "error"
    },
    "data": null
    }
  ```

***Delete Book***
- URL: /books/:id
- Method: DELETE
- Response:

``` json
{
  "meta": {
    "message": "Successfully deleted book",
    "code": 200,
    "status": "success"
  },
  "data": null
}
```
- Response Error
  - Book by Id Not Found
  ```json
    {
    "meta": {
        "message": "Data Not Found",
        "code": 404,
        "status": "error"
    },
    "data": null
    }
  ```

## Data Model

***Book***
- ID: uint
- Title: string
- Author: string
- Price: int (formatted as "Rp " with thousand separators)
- Quantity: int
- PageCount: int
- Description: string
- CreatedAt: time.Time (formatted as "DD MMMM YYYY HH")
- UpdatedAt: time.Time (formatted as "DD MMMM YYYY HH")

    


