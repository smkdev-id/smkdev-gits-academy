# Todo API

Todo API adalah proyek yang dibangun menggunakan Golang dan framework Echo untuk mengelola daftar tugas (todo list). API ini menyediakan berbagai endpoint untuk membuat, membaca, memperbarui, dan menghapus tugas.

---

## Authors

- [@Tri Krama](https://www.github.com/trikrama)

---

## Fitur ✨

- CRUD Todo:
  - **GET /api/v1/todo**: Mendapatkan semua todo.
  - **GET /api/v1/todo/:id**: Mendapatkan todo berdasarkan ID.
  - **POST /api/v1/todo**: Membuat todo baru.
  - **PUT /api/v1/todo/:id**: Memperbarui todo.
  - **DELETE /api/v1/todo/:id**: Menghapus todo.
  
---

## Instalasi 👨🏻‍💻

Proyek ini dibangun menggunakan [Go](https://go.dev/dl/) 1.13+.

1. Clone repository

```sh
- git clone https://github.com/trikrama/todos.git
- cd todos
- go mod tidy
- go run main.go
```

---


## API Reference

#### Get All Todo

```http
  GET /api/v1/todo
```


#### Get Todo By Id

```http
  GET /api/v1/todo/:id
```

#### Create Todo

```http
  POST /api/v1/todo
```

```sh
  -H 'Content-Type: application/json' \
  -d '{
      "title": "Learn Golang",
      "description": "Learn how to use Echo with Golang",
      "completed": false
      }'
```

#### Update Todo

```http
  PUT /api/v1/todo/:id
```

```sh
  -H 'Content-Type: application/json' \
  -d '{
      "title": "Learn Golang",
      "description": "Learn how to use Echo with Golang",
      "completed": false
      }'
```

#### Delete Todo

```http
  DELETE /api/v1/todo/:id
```

---


## Tech Stack

**Database:** Sqlite

**Framework:** Echo Golang



---


## License

[MIT](https://choosealicense.com/licenses/mit/)


Proyek ini dilisensikan di bawah MIT License.

Dilarang menjual atau mendistribusikan perangkat lunak ini untuk tujuan komersial tanpa izin tertulis dari kami.




