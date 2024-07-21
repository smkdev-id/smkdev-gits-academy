
# Menjalankan Project Golang-CRUD

Berikut adalah cara menjalankan project golang



## Command Line Terminal

Instal Depedensi

```bash
go get github.com/labstack/echo/v4
go get modernc.org/sqlite

```
Jalankan Program

```bash
go run .

```




## Demo

Cek di Postman:

Get: http://localhost:8080/users = Menampilkan data yang berada pada database

Post: http://localhost:8080/users = Masukan data raw

```bash
{
    "name": "John Doe",
    "email": "john.doe@example.com",
    "password": "securepassword"
}

```

Put: http://localhost:8080/users/1 = Update Data no 1
```bash
{
    "name": "Jhan Dha",
    "email": "jhandha@example.com",
    "password": "admin123kominfo"
}

```
Delete: http://localhost:8080/users/1 = Delete Data no 1




