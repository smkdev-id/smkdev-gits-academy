package main

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	_ "modernc.org/sqlite"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var db *sql.DB

func main() {
	var err error
	// Membuka koneksi ke database SQLite
	db, err = sql.Open("sqlite", "smkdev")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	createTable()

	e := echo.New()
	// Endpoint CRUD
	e.GET("/users", getUsers)
	e.POST("/users", createUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)

	e.Logger.Fatal(e.Start(":8080"))
}

// createTable membuat tabel users jika belum ada
func createTable() {
	createTableSQL := `CREATE TABLE IF NOT EXISTS users (
        "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        "name" TEXT,
        "email" TEXT,
        "password" TEXT
    );`
	_, err := db.Exec(createTableSQL)
	if err != nil {
		panic(err)
	}
}

// getUsers mengambil semua pengguna dari database
func getUsers(c echo.Context) error {
	rows, err := db.Query("SELECT id, name, email, password FROM users")
	if err != nil {
		return err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password); err != nil {
			return err
		}
		users = append(users, user)
	}
	return c.JSON(http.StatusOK, users)
}

// createUser membuat pengguna baru di database
func createUser(c echo.Context) error {
	user := new(User)
	if err := c.Bind(user); err != nil {
		return err
	}

	stmt, err := db.Prepare("INSERT INTO users (name, email, password) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	res, err := stmt.Exec(user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	user.ID = int(id)
	return c.JSON(http.StatusOK, user)
}

// updateUser mengubah data pengguna di database berdasarkan ID
func updateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	user := new(User)
	if err := c.Bind(user); err != nil {
		return err
	}

	stmt, err := db.Prepare("UPDATE users SET name = ?, email = ?, password = ? WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(user.Name, user.Email, user.Password, id)
	if err != nil {
		return err
	}
	user.ID = id
	return c.JSON(http.StatusOK, user)
}

// deleteUser menghapus pengguna dari database berdasarkan ID
func deleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	stmt, err := db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}
