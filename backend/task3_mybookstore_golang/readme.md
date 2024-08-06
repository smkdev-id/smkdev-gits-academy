# Bookstore Management API

## Tech Stack

- Go (Golang)
- Gorilla Mux
- GORM
- SQLite
- GoDotEnv

## Project Overview

Bookstore Management API is a RESTful service for managing books, transactions, and customer information in a bookstore. It includes authentication, admin routes for managing books and transactions, and customer routes for managing baskets and transactions.

## Project Structure

task3_mybookstore_golang/
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
│ └── utils/
│ └── utils.go
└── go.mod
└── go.sum
└── books.db

## Steps to Run the Project

1. Clone the repository:

   ```sh
   git clone https://github.com/mohrizkyputraajiutama/task3_mybookstore_golang.git
   cd task3_mybookstore_golang
   ```

2. Install dependencies:

   ```sh
   go mod tidy
   ```

3. Create a `.env` file and add the following:

   ```
   DB_NAME=books.db
   PORT=8080
   ```

4. Run the project:
   ```sh
   go run cmd/main.go
   ```

## API

### Authentication

**Headers**

- `Content-Type: application/json`
- `Authorization: Bearer <token>`

### Admin Routes

#### Books Management

- **Create Book**

  - **Endpoint:** `/book/`
  - **Method:** `POST`
  - **Description:** Adds a new book.

- **Get All Books**

  - **Endpoint:** `/book/`
  - **Method:** `GET`
  - **Description:** Retrieves all books.

- **Get Book By ID**

  - **Endpoint:** `/book/{bookId}`
  - **Method:** `GET`
  - **Description:** Retrieves a book by ID.

- **Update Book**

  - **Endpoint:** `/book/{bookId}`
  - **Method:** `PUT`
  - **Description:** Updates a book by ID.

- **Delete Book**
  - **Endpoint:** `/book/{bookId}`
  - **Method:** `DELETE`
  - **Description:** Deletes a book by ID.

#### Transaction Management

- **Create Transaction**

  - **Endpoint:** `/transaction/`
  - **Method:** `POST`
  - **Description:** Creates a new transaction.

- **Get All Transactions**

  - **Endpoint:** `/transaction/`
  - **Method:** `GET`
  - **Description:** Retrieves all transactions.

- **Get Transaction By ID**

  - **Endpoint:** `/transaction/{transactionId}`
  - **Method:** `GET`
  - **Description:** Retrieves a transaction by ID.

- **Update Transaction**

  - **Endpoint:** `/transaction/{transactionId}`
  - **Method:** `PUT`
  - **Description:** Updates a transaction by ID.

- **Delete Transaction**
  - **Endpoint:** `/transaction/{transactionId}`
  - **Method:** `DELETE`
  - **Description:** Deletes a transaction by ID.

### Customer Routes

#### Books

- **Get All Books**
  - **Endpoint:** `/books/`
  - **Method:** `GET`
  - **Description:** Retrieves all available books for customers.

#### Basket Management

- **Add to Basket**

  - **Endpoint:** `/basket/`
  - **Method:** `POST`
  - **Description:** Adds a book to the customer's basket.

- **Get Basket**

  - **Endpoint:** `/basket/`
  - **Method:** `GET`
  - **Description:** Retrieves the current basket for the customer.

- **Remove from Basket**
  - **Endpoint:** `/basket/{bookId}`
  - **Method:** `DELETE`
  - **Description:** Removes a book from the customer's basket.

#### Transaction Management

- **Create Transaction**

  - **Endpoint:** `/customer/transaction/`
  - **Method:** `POST`
  - **Description:** Creates a new transaction for the customer.

- **Get Transactions**
  - **Endpoint:** `/customer/transactions/`
  - **Method:** `GET`
  - **Description:** Retrieves all transactions for the customer.

## How to Use

### Authentication

#### Headers

- `Content-Type: application/json`

#### Register

- **Endpoint:** `/auth/register`
- **Method:** `POST`
- **Description:** Registers a new user.

#### Login

- **Endpoint:** `/auth/login`
- **Method:** `POST`
- **Description:** Logs in an existing user and returns a JWT token.

### Admin Endpoints

Use the following headers for all admin endpoints:

- `Content-Type: application/json`
- `Authorization: Bearer <admin_token>`

- **Create Book**

  - Example:
    ```json
    {
      "title": "Book Title",
      "author": "Author Name",
      "price": 19.99,
      "stock": 10
    }
    ```

- **Update Book**
  - Example:
    ```json
    {
      "title": "Updated Book Title",
      "author": "Updated Author Name",
      "price": 29.99,
      "stock": 5
    }
    ```

### Customer Endpoints

Use the following headers for all customer endpoints:

- `Content-Type: application/json`
- `Authorization: Bearer <customer_token>`

- **Add to Basket**

  - Example:
    ```json
    {
      "bookId": "1",
      "quantity": 2
    }
    ```

- **Create Transaction**
  - Example:
    ```json
    {
      "paymentMethod": "Credit Card",
      "address": "Customer Address"
    }
    ```
