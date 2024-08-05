# Book Sales Management API

Book Sales Management API adalah aplikasi API sederhana untuk mengelola penjualan buku. API ini dibangun menggunakan Golang, GORM sebagai ORM, dan SQLite sebagai basis data. API ini menyediakan endpoint untuk melakukan operasi CRUD (Create, Read, Update, Delete) pada data buku.

## Struktur Proyek
Berikut adalah struktur direktori proyek Book Sales Management API:
- config
  - app.go
- controllers
  - book_controllers.go
- models
  - book.go
- routes
  - book_routes.go
- utils
  - utils.go
- main.go


## Penjelasan File dan Direktori

- **config/app.go:** File konfigurasi aplikasi, termasuk inisialisasi database.
- **controllers/book_controllers.go:** Controller untuk mengelola request terkait buku.
- **models/book.go:** Model untuk entitas buku, termasuk definisi schema dan fungsi terkait.
- **routes/book_routes.go:** Definisi routing untuk endpoint terkait buku.
- **utils/utils.go:** Berisi fungsi utilitas umum yang digunakan di seluruh aplikasi.
- **main.go:** File utama untuk menjalankan aplikasi.

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
  {
    "data": [
        {
            "ID": 1,
            "CreatedAt": "2024-08-05T08:19:35.3685923+07:00",
            "UpdatedAt": "2024-08-05T08:27:52.8331292+07:00",
            "DeletedAt": null,
            "title": "Laskar Pelangi",
            "author": "Andrea Hirata",
            "price": 50,
            "publisher": "Bentang Pustaka",
            "genre": "Fiksi",
            "page_count": 529,
            "language": "Indonesian"
        },
        {
            "ID": 2,
            "CreatedAt": "2024-08-05T08:28:27.3463732+07:00",
            "UpdatedAt": "2024-08-05T08:28:27.3463732+07:00",
            "DeletedAt": null,
            "title": "Bumi Manusia",
            "author": "Pramoedya Ananta Toer",
            "price": 65.5,
            "publisher": "Lentera Dipantara",
            "genre": "Fiksi Sejarah",
            "page_count": 535,
            "language": "Indonesian"
        }
    ],
    "message": "Data retrieved successfully"
  }
  ```

### 2. Create Book
- **URL:** `/book/`
- **Method:** `POST`
- **Description:** Membuat buku baru.
- **Request Body:**
  ```json
  {
    "title": "Laskar Pelangi",
    "author": "Andrea Hirata",
    "price": 50.00,
    "publisher": "Bentang Pustaka",
    "genre": "Fiksi",
    "page_count": 529,
    "language": "Indonesian"
  }
  ```
- **Response:**
  ```json
  {
    "data": {
       "ID": 2,
        "CreatedAt": "2024-08-05T08:28:27.3463732+07:00",
        "UpdatedAt": "2024-08-05T08:28:27.3463732+07:00",
        "DeletedAt": null,
        "title": "Bumi Manusia",
        "author": "Pramoedya Ananta Toer",
        "price": 65.5,
        "publisher": "Lentera Dipantara",
        "genre": "Fiksi Sejarah",
        "page_count": 535,
        "language": "Indonesian"
    },
    "message": "Book created successfully"
  }
  ```

### 3. Get Book by ID
- **URL:** `/book/{bookId}`
- **Method:** `GET`
- **Description:** Mendapatkan data buku berdasarkan ID.
- **Path parameter:** `bookId = ID dari buku`
- **Response:**
  ```json
  {
    "data": {
        "ID": 2,
        "CreatedAt": "2024-08-05T08:28:27.3463732+07:00",
        "UpdatedAt": "2024-08-05T08:28:27.3463732+07:00",
        "DeletedAt": null,
        "title": "Bumi Manusia",
        "author": "Pramoedya Ananta Toer",
        "price": 65.5,
        "publisher": "Lentera Dipantara",
        "genre": "Fiksi Sejarah",
        "page_count": 535,
        "language": "Indonesian"
    },
    "message": "Book retrieved successfully"
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
    "title": "Negeri 5 Menara",
    "author": "Ahmad Fuadi",
    "price": 45.30,
    "publisher": "Gramedia Pustaka Utama",
    "genre": "Fiksi Inspiratif",
    "page_count": 415,
    "language": "Indonesian"
  }
  ```
- **Response:**
  ```json
  {
    "data": {
        "ID": 2,
        "CreatedAt": "2024-08-05T08:28:27.3463732+07:00",
        "UpdatedAt": "2024-08-05T08:30:48.4031316+07:00",
        "DeletedAt": null,
        "title": "Negeri 5 Menara",
        "author": "Ahmad Fuadi",
        "price": 45.3,
        "publisher": "Gramedia Pustaka Utama",
        "genre": "Fiksi Inspiratif",
        "page_count": 415,
        "language": "Indonesian"
    },
    "message": "Book updated successfully"
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
    "data": {
        "ID": 2,
        "CreatedAt": "2024-08-05T08:28:27.3463732+07:00",
        "UpdatedAt": "2024-08-05T08:30:48.4031316+07:00",
        "DeletedAt": "2024-08-05T08:31:38.8702482+07:00",
        "title": "Negeri 5 Menara",
        "author": "Ahmad Fuadi",
        "price": 45.3,
        "publisher": "Gramedia Pustaka Utama",
        "genre": "Fiksi Inspiratif",
        "page_count": 415,
        "language": "Indonesian"
    },
    "message": "Book deleted successfully"
  }
  ```

## Cara Menjalankan Proyek

### 1. **Kloning Repository**
```sh
git clone https://github.com/smkdev-id/smkdev-gits-academy.git
git checkout muhammad-albert-nur-agathon
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