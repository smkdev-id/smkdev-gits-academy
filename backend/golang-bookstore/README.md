# Bookstore API

## Deskripsi Proyek

Bookstore API adalah aplikasi backend yang dirancang untuk mengelola penjualan buku. API ini dibangun menggunakan bahasa pemrograman Go, framework Gin, dan GORM sebagai ORM dengan SQLite sebagai database. Aplikasi ini menyediakan endpoint CRUD (Create, Read, Update, Delete) untuk mengelola data buku.

### Fitur Utama

- **Create**: Membuat buku baru di database.
- **Read**: Mendapatkan detail buku atau daftar buku.
- **Update**: Memperbarui informasi buku berdasarkan ID.
- **Delete**: Menghapus buku dari database berdasarkan ID.

## Instalasi

### Prasyarat

1. **Go**: Pastikan Go versi 1.20 atau lebih baru sudah terpasang di sistem kamu. Kamu bisa mengunduhnya dari [situs resmi Go](https://golang.org/dl/).
2. **SQLite**: SQLite biasanya sudah terintegrasi dalam GORM, tetapi pastikan SQLite sudah tersedia di sistem kamu jika dibutuhkan.

### Langkah-langkah Instalasi

1. **Clone Repositori**

   Pertama, clone repositori ini ke sistem lokal kamu:

   ```bash
   git clone https://github.com/smkdev-id/smkdev-gits-academy.git
   cd backend/golang-bookstore

## Instalasi Dependensi
Install dependensi yang dibutuhkan dengan perintah :

```bash
go mod tidy
```

## Konfigurasi Database
Konfigurasi database SQLite dapat dilakukan di file pkg/config/app.go. Pastikan konfigurasi sudah sesuai dengan kebutuhan kamu.

## Jalankan Aplikasi
```bash 
go run cmd/main.go
```
Aplikasi akan berjalan di http://localhost:3000.

## Dokumentasi API
Dokumentasi API tersedia di folder root projek dengan nama file api-spec.json