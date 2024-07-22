
## API Reference
Base URL: `http://localhost:8080`

## Endpoints

### Create User

Creates a new user.

- **URL:** `/users`
- **Method:** `POST`
- **Request Body:**
  ```json
  {
    "name": "John Doe",
    "email": "john@example.com",
    "password": "securepassword"
  }

### Read user

Read a new user.

- **URL:** `/users`
- **Method:** `GET`

### Update User

Update a new user.

- **URL:** `/users/:id`
- **Method:** `PUT`
- **Request Body:**
  ```json
    {
    "name": "John Updated",
    "email": "john.updated@example.com"
    }
### Delete User

Delete a new user.

- **URL:** `/users/:id`
- **Method:** `DELETE`
