# Backend Book Store

## Description

Backend Book Store is a CRUD application for managing book data in a bookstore. It is built using Go and utilizes Gin for handling HTTP requests, GORM for ORM, and Redis for caching.

## Technologies Used

- Go 1.22.5
- Gin v1.10.0
- GORM v1.25.11
- Redis v8.11.5

## Endpoints

### GET /books

- Description: Get all books
- Response: JSON array of book objects

### POST /books

- Description: Add a new book
- Request Body: JSON object with book details
- Response: JSON object of the added book

### PUT /books/{id}

- Description: Update a book by ID
- Request Body: JSON object with updated book details
- Response: JSON object of the updated book

### DELETE /books/{id}

- Description: Delete a book by ID
- Response: Success message

### POST /register

- Description: Register new user
- Response: JSON object with user info and jwt token

### POST /login

- Description: Login user
- Response: JSON message with token

### POST /logout

- Description: Logout user
- Response: JSON message "Logout Success"

### POST /assign-role/{userId}

- Description: Assign role to user
- Response: JSON message "Role assigned to user"

### GET /role

- Description: Get all role
- Response: JSON object with role data

### POST /role

- Description: Create role
- Response: JSON object with role data

## License

This project is licensed under the MIT License. See the LICENSE.md file for details.
