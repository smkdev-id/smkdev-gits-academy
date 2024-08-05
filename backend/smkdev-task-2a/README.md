# Simple User Management API

Proyek ini adalah sebuah API sederhana untuk manajemen pengguna menggunakan Go, Echo framework, dan SQLite database.

## Fitur

- Membuat pengguna baru
- Memperbarui informasi pengguna
- Menghapus pengguna
- Mendapatkan daftar semua pengguna

## Persyaratan

- Go (versi 1.16 atau lebih baru)
- SQLite

## Teknologi yang Digunakan

- [Go](https://golang.org/)
- [Echo Framework](https://echo.labstack.com/)
- [SQLite](https://www.sqlite.org/)
- [modernc.org/sqlite](https://pkg.go.dev/modernc.org/sqlite)

## Kontribusi

Kontribusi selalu diterima. Silakan buka issue atau submit pull request jika Anda ingin berkontribusi.

## Lisensi

[MIT License](https://opensource.org/licenses/MIT)

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
