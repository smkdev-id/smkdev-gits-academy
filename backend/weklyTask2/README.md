# Weekly Task Todolist API

Proyek ini adalah sebuah REST API sederhana untuk mengelola todo list harian menggunakan Golang, Echo framework, dan SQLite sebagai database.

## Daftar Isi

- [Instalasi](#instalasi)
- [Menjalankan Aplikasi](#menjalankan-aplikasi)
- [Endpoint API](#endpoint-api)
  - [Get Todo List](#get-todo-list)
  - [Post Todo List](#post-todo-list)
  - [Patch Todo List](#patch-todo-list)
  - [Delete Todo List](#delete-todo-list)
- [Struktur Proyek](#struktur-proyek)

## Instalasi

Any additional information goes here

1. Clone repository ini:
   ```bash
   git clone https://github.com/smkdev-id/smkdev-gits-academy.git
   cd weklyTask2
   ```
   2.Instal dependencis :

```bash
go get github.com/labstack/echo/v4
go get github.com/mattn/go-sqlite3
```

Atau

```bash
go mod tidy
```

## Menjalankan Aplikasi Todo List

1. Jalankan Server :

```bash
 go run main.go
```

2. Server akan berjalan di port `http://localhost:3030`

## Endpoint API Todo List

#### Get Todo List

- **Endpoint:** `GET /todos`
- **Deskripsi:** Mengambil semua todo list yang ada.
- **Contoh Permintaan Di Postman:**
  1. Buka Postman.
  2. Buat permintaan baru dengan metode `GET`.
  3. Masukkan URL: `http://localhost:3030/todos`.
  4. Klik tombol `Send`.

#### Post Todo List

- **Endpoint:** `POST /todos`
- **Deskripsi:** Membuat todo list baru
- **Contoh Permintaan Di Postman:**
  1. Buka Postman.
  2. Buat Permintaan baru dengan metode `POST`.
  3. Masukkan URL: `http://localhost:3030/todos`.
  4. Pilih tab `Body`.
  5. Pilih `raw` dan `json` sebagai type
  6. Masukkan data berikut di dalam body:
  ```json
  {
    "todo": "Besok Kelas Gits Academy",
    "completed": false
  }
  ```
  7. Klik tombol `Send`

#### Patch Todo List

- **Endpoint:** `PATCH /todos/:id`
- **Deskripsi:** Mengedit todo list yang ada. Bisa mengganti todo saja, Bisa juga hanya status completednya saja, atau Bisa keduanya
- **Contoh Permintaan Di Postman:**
  1. Buka Postman.
  2. Buat Permintaan baru dengan metode `PATCH`.
  3. Masukkan URL: `http://localhost:3030/todos/2`.
  4. Pilih tab `Body`.
  5. Pilih `raw` dan `json` sebagai type.
  6. Masukkan data berikut di dalam body:
  ```json
  {
    "completed": true
  }
  ```
  7. Klik tombol `Send`

#### Delete Todo List

- **Endpoint:** `DELETE /todos/:id`
- **Deskripsi:** Menghapus todo list berdasarkan ID.
- **Contoh Permintaan Di Postman:**
  1. Buka Postman.
  2. Buat permintaan baru dengan metode `DELETE`.
  3. Masukkan URL: `http://localhost:3030/todos/2`.
  4. Klik tombol `Send`.

## Struktur Proyek

```bash
.
├── main.go               # Entry point aplikasi
├── go.sum
├── go.mod
├── utils                 # Folder untuk utilitas umum
│   └── response.go
├── services              # Folder untuk service layer
│   └── todoService.go
├── routes                # Folder untuk routing
│   └── routes.go
├── models                # Folder untuk model
│   └── todo.go
├── controllers           # Folder untuk handler HTTP (controller)
│   └── todoController.go
└── config                # Folder untuk konfigurasi aplikasi dan database
    └── config.go

```
