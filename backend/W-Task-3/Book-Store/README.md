
# Book Store API

## Deskripsi

The Book Store API is a RESTful service built with Go (Golang) that provides CRUD operations for managing a book inventory. It uses GORM for database interactions with SQLite as the storage backend and Gorilla Mux for routing. This API allows users to perform operations such as retrieving all books, fetching a book by ID, creating new books, updating existing books, and deleting books from the database. The application is designed to facilitate efficient management of book data through a simple and intuitive API.





## Prerequisites
Before you begin, ensure you have the following installed:
- [Go](https://golang.org/dl/)
- [TDM-GCC](https://jmeubank.github.io/tdm-gcc/) or MinGW for C compiler on Windows.
- [SQLite](https://www.sqlite.org/)
- [Postman](https://www.postman.com/downloads/) for API testing.
- [GORM and Gorilla Mux Dependencies] The application uses GORM for ORM (Object-Relational Mapping) and Gorilla Mux for routing. These dependencies will be automatically managed and installed using Go modules when you run go mod tidy.
## Installation

Follow these steps to set up and run the Book Store API

1. Clone the Repository
    Open your terminal and clone the repository using Git:
    ```sh
    git clone https://github.com/slthn2025/Book-Store.git 
    cd Book-Store

    ```

2. Install Dependencies:

    ```sh
    go mod tidy

    ```

3. Run the Application:

    ```sh
    go run cmd/main.go

    ```

4. Verify the Server:

Open your web browser or API client (e.g., Postman) and navigate to   http://localhost:8080. You can test the API endpoints using the provided URLs in the API Endpoints section.

The application will run on `http://localhost:8080`.

## Accessing Endpoints with Postman

### GET /books
- Method: GET
- URL: `http://localhost:8080/books`
- Description: Retrieves a list of all books.

### POST /items
- Method: POST
- URL: `http://localhost:8080/books`
- Body:
    ```json
    {
        "name": "ItemName",
        "value": "ItemValue"
    }
    ```
- Description: Adds a new books.

### PUT /items/:id
- Method: PUT
- URL: `http://localhost:8080/books/{id}`
- Body:
    ```json
    {
        "name": "UpdatedItemName",
        "value": "UpdatedItemValue"
    }
    ```
- Description: Updates an books by ID.

### DELETE /items/:id
- Method: DELETE
- URL: `http://localhost:8080/books/{id}`
- Description: Deletes an books by ID.

    