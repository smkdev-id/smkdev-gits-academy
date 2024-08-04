# Book Sales Management API

Book Sales Management API adalah aplikasi API sederhana untuk mengelola penjualan buku. API ini dibangun menggunakan Golang, GORM sebagai ORM, dan SQLite sebagai basis data. API ini menyediakan endpoint untuk melakukan operasi CRUD (Create, Read, Update, Delete) pada data buku.

## Struktur Proyek


## Konfigurasi

#### 1. **Database:** SQLite (`bookstore.db`).
#### 2. **ORM:** GORM.
#### 3. **Server:** Gin.

## Endpoint API
### 1. Get All Books
- **URL:** `/book/`
- **Method:** `GET`
- **Description:** Mendapatkan daftar semua buku.
- **Response:**
  ```json
  [
    {
      "ID": 1,
      "CreatedAt": "2024-08-03T10:00:00Z",
      "UpdatedAt": "2024-08-03T10:00:00Z",
      "DeletedAt": null,
      "title": "The Great Gatsby",
      "author": "F. Scott Fitzgerald",
      "price": 10.99
    }
  ]
  ```

### 2. Create Book
- **URL:** `/book/`
- **Method:** `POST`
- **Description:** Membuat buku baru.
- **Request Body:**
  ```json
  {
    "title": "To Kill a Mockingbird",
    "author": "Harper Lee",
    "price": 7.99
  }
  ```
- **Response:**
  ```json
  [
    {
      "ID": 1,
      "CreatedAt": "2024-08-03T10:00:00Z",
      "UpdatedAt": "2024-08-03T10:00:00Z",
      "DeletedAt": null,
      "title": "The Great Gatsby",
      "author": "F. Scott Fitzgerald",
      "price": 10.99
    }
  ]
  ```

### 3. Get Book by ID
- **URL:** `/book/{bookId}`
- **Method:** `GET`
- **Description:** Mendapatkan data buku berdasarkan ID.
- **Path parameter:** `bookId = ID dari buku`
- **Response:**
  ```json
  {
    "ID": 1,
    "CreatedAt": "2024-08-03T10:00:00Z",
    "UpdatedAt": "2024-08-03T10:00:00Z",
    "DeletedAt": null,
    "title": "The Great Gatsby",
    "author": "F. Scott Fitzgerald",
    "price": 10.99
  }
  ```

### 4. Update Book
- **URL:** `/book/{bookId}`
- **Method:** `PUT`
- **Description:** Memperbarui data buku.
- **Path parameter:** `bookId = ID dari buku`
- **Request Body:**
  ```json
    {
      "title": "To Kill a Mockingbird",
      "author": "Harper Lee",
      "price": 7.99
    }
  ```
- **Response:**
  ```json
  {
    "ID": 1,
    "CreatedAt": "2024-08-03T10:00:00Z",
    "UpdatedAt": "2024-08-03T10:00:00Z",
    "DeletedAt": null,
    "title": "The Great Gatsby",
    "author": "F. Scott Fitzgerald",
    "price": 10.99
  }
  ```

### 5. Delete Book
- **URL:** `/book/{bookId}`
- **Method:** `DELETE`
- **Description:** Menghapus data buku.
- **Path parameter:** `bookId = ID dari buku`
- **Response:**
  ```json
  {
    "message": "Book deleted",
  }
  ```

## Cara Menjalankan Proyek

### 1. **Kloning Repository**
```sh
git clone <repository-url>
cd <repository-directory>
```

### 2. **Instalasi Dependensi**
```sh
go mod tidy
```

### 3. **Menjalankan Server**
```sh
go run main.go
```

## Penggunaan
Gunakan alat seperti Postman atau curl untuk mengirim permintaan HTTP ke endpoint API.

## Catatan
- **Mode**: Secara default, Gin berjalan dalam "debug" mode. Untuk menjalankan dalam "release" mode, set environment variable GIN_MODE menjadi release atau gunakan gin.SetMode(gin.ReleaseMode) dalam kode.

- **Proxies**: Jika aplikasi ini berjalan di belakang proxy atau load balancer, pastikan untuk mengkonfigurasi proxy yang dipercaya menggunakan router.SetTrustedProxies.


### Penjelasan

- **Struktur Proyek:** Menjelaskan bagaimana struktur folder dan file diatur dalam proyek.
- **Konfigurasi:** Menyebutkan teknologi yang digunakan.
- **Endpoint API:** Menjelaskan setiap endpoint yang tersedia, metode HTTP yang digunakan, deskripsi, parameter, dan contoh respons.
- **Cara Menjalankan Proyek:** Langkah-langkah untuk menjalankan server.
- **Penggunaan:** Saran untuk menggunakan alat seperti Postman.
- **Catatan:** Informasi penting seperti pengaturan mode aplikasi dan proxy.

Dokumentasi ini bisa dikembangkan lebih lanjut sesuai dengan kebutuhan proyek dan informasi tambahan yang perlu disertakan.