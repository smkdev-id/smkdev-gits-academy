# Bookstore API

API ini menyediakan beberapa endpoint untuk mengelola data buku di sebuah toko buku. Anda dapat menambahkan, memperbarui, menghapus, dan mengambil data buku melalui API ini.

## Daftar Endpoint

### 1. Fetch All Book

**Endpoint:** `/book`  
**Method:** `GET`  
**Deskripsi:** Mengambil semua data buku yang ada di database.

### 2. Fetch Book By id

**Endpoint:** `/book/:id`  
**Method:** `GET`  
**Deskripsi:** Mengambil data buku yang ada di database berdasarkan id.

### 3. Create A New Book

**Endpoint:** `/book`  
**Method:** `POST`  
**Deskripsi:** Menambahkan buku baru ke dalam database.

**Request Body:**
```json
{
    "nama_buku": "Filosofi Teras",
    "nama_penulis": "Henry Manampiring",
    "penerbit": "PT Kompas Media Nusantara",
    "stock": 100,
    "harga_buku": 100000
}
```

### 4. Update A Book

**Endpoint:** `/book/:id`  
**Method:** `PUT`  
**Deskripsi:** Mengubah buku yang sudah ada di database. Tinggal cantumkan saja properti buku yang mana yang akan di ubah

Contoh: 

**Request Body:**
```json
{
    "nama_buku": "Belajar Ngoding Untuk Bayi",
    "nama_penulis": "Hilman",
    "penerbit": "PT Kompas Media Nusantara",
    "stock": 100,
    "harga_buku": 100000
}
```

Ataupun hanya salah satu properti saja

Contoh saya ingin mengupdate stock buku nya

**Request Body:**
```json
{
    "stock": 100
}
```

### 5. Delete A Book By ID

**Endpoint:** `/book/:id`  
**Method:** `DELETE`  
**Deskripsi:** Menghapus data buku yang ada di database berdasarkan id.

Jika ingin mencoba nya langsung silahkan di clone repository ini dan import postmant collection yang tersedia.

Made With Golang.
