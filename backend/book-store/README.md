# Book Store API

Book Store API adalah proyek yang dibangun menggunakan Golang dan framework Echo untuk mengelola daftar tugas (book store). API ini menyediakan berbagai endpoint untuk membuat, membaca, memperbarui, dan menghapus tugas.

---

## Authors

- [@Rayhan Alsauqi](https://github.com/RayhanAsauqi)

---

## Fitur ✨

- CRUD Book Store:
  - **GET /api/v1/books**: Mendapatkan semua books.
  - **GET /api/v1/books/:bookId**: Mendapatkan book berdasarkan ID.
  - **POST /api/v1/books**: Membuat book baru.
  - **PATCH /api/v1/books/:bookId**: Memperbarui book.
  - **DELETE /api/v1/books/:bookId**: Menghapus book.

---

## Instalasi 👨🏻‍💻

Proyek ini dibangun menggunakan [Go](https://go.dev/dl/) 1.13+.

1. Clone repository

```sh
- git clone https://github.com/smkdev-id/smkdev-gits-academy/tree/Rayhan-Alsauqi
```

2. Masuk Ke Direktori

```sh
- cd book-store
```

3. Install Package

```sh
- go mod tidy
```

4. Run Program

```sh
- go run cmd/main.go
```

---

## API Reference

#### Get All Books

```http
  GET /api/v1/books
```

#### Get Book By Id

```http
  GET /api/v1/books/:bookId
```

#### Create Book

```http
  POST /api/v1/books
```

```sh
  -H 'Content-Type: application/json' \
  -d '{
      "title": "Contoh Buku1",
      "author": "Nama Penulis1",
      "publisher": "Nama Penerbit1",
      "description": "Deskripsi singkat tentang buku ini1."
      }'
```

#### Update Todo

```http
  PUT /api/v1/todo/:id
```

```sh
  -H 'Content-Type: application/json' \
  -d '{
      "title": "Buku1",
      "author": "Nama Penulis1"
      }'
```

#### Delete book

```http
  DELETE /api/v1/books/:bookId
```

---

## Tech Stack

**Database:** SQLite

**Framework:** Echo Golang

---

## License

[MIT](https://choosealicense.com/licenses/mit/)

Proyek ini dilisensikan di bawah MIT License.

Dilarang menjual atau mendistribusikan perangkat lunak ini untuk tujuan komersial tanpa izin tertulis dari kami.
