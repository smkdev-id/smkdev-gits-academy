# task_week2_golang

Project ini adalah implementasi dari CRUD (Create, Read, Update, Delete) endpoint menggunakan bahasa pemrograman Go (Golang) dan SQLite sebagai database. Endpoint ini dapat diuji menggunakan Postman.

## Tools yang Digunakan

1. **Go (Golang)**

   - Bahasa pemrograman utama untuk pengembangan aplikasi ini.
   - [Download Go](https://golang.org/dl/)

2. **SQLite**

   - Database yang digunakan untuk menyimpan data.
   - [Download SQLite](https://www.sqlite.org/download.html)

3. **Postman**
   - Alat untuk menguji API endpoint.
   - [Download Postman](https://www.postman.com/downloads/)

## Cara Menggunakan

### Persiapan

1. **Instal Golang:**

   - Pastikan Go sudah terinstal di sistem Anda. Jika belum, unduh dan instal dari [tautan ini](https://golang.org/dl/).

2. **Instal SQLite:**

   - Pastikan SQLite sudah terinstal di sistem Anda. Jika belum, unduh dan instal dari [tautan ini](https://www.sqlite.org/download.html).

3. **Instal Postman:**
   - Pastikan Postman sudah terinstal di sistem Anda. Jika belum, unduh dan instal dari [tautan ini](https://www.postman.com/downloads/).

### Menjalankan Proyek

1. **Clone Repository:**

   - Clone repository ini ke dalam direktori lokal Anda.
     ```bash
     git clone <URL_REPOSITORY>
     cd task_week2_golang
     ```

2. **Inisialisasi Modul Go:**

   - Inisialisasi dan unduh dependensi yang diperlukan.
     ```bash
     go mod init task_week2_golang
     go mod tidy
     ```

3. **Jalankan Aplikasi:**
   - Jalankan aplikasi menggunakan perintah berikut:
     ```bash
     go run main.go
     ```

### Menggunakan Postman untuk Menguji Endpoint

1. **Create (POST /todos):**

   - Metode: `POST`
   - URL: `http://localhost:8080/todos`
   - Body (raw JSON):
     ```json
     {
       "title": "Learn Golang",
       "description": "Complete the task_week2_golang project",
       "completed": false
     }
     ```

2. **Read (GET /todos):**

   - Metode: `GET`
   - URL: `http://localhost:8080/todos`

3. **Update (PUT /todos/{id}):**

   - Metode: `PUT`
   - URL: `http://localhost:8080/todos/{id}` (gantilah `{id}` dengan ID todo yang ingin Anda perbarui)
   - Body (raw JSON):
     ```json
     {
       "title": "Learn Golang",
       "description": "Complete the task_week2_golang project",
       "completed": true
     }
     ```

4. **Delete (DELETE /todos/{id}):**
   - Metode: `DELETE`
   - URL: `http://localhost:8080/todos/{id}` (gantilah `{id}` dengan ID todo yang ingin Anda hapus)

## Bahasa yang Digunakan

- **Go (Golang):**

  - Golang adalah bahasa pemrograman yang dikembangkan oleh Google. Bahasa ini dirancang untuk pengembangan aplikasi yang efisien dan efektif, terutama untuk aplikasi backend. (https://golang.org/doc/)

- **SQL (SQLite):**
  - SQLite adalah sistem manajemen database relasional yang ringan. SQLite sering digunakan untuk aplikasi kecil hingga menengah karena kesederhanaannya. (https://www.sqlite.org/docs.html)

## Struktur Proyek

## Struktur Proyek

- **main.go**
  - File utama untuk menjalankan aplikasi. Menyediakan konfigurasi router dan endpoint CRUD.
- **models/todo.go**
  - Mendefinisikan model data untuk aplikasi ini, termasuk struktur data untuk `todos`.
- **handlers.go**

  - Berisi fungsi-fungsi untuk menangani permintaan API untuk operasi CRUD.

- **database/database.go**

  - Mengatur koneksi ke database SQLite dan operasi dasar database.

- **go.mod**

  - File modul Go yang mengelola dependensi proyek.

- **go.sum**
  - File yang berisi checksum untuk memastikan integritas dependensi Go.
- **todos.db**
  - File untuk mengintegrasikan ke database sqlite
