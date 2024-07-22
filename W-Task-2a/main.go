package main

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("sqlite3", "./user.db")
	if err != nil {
		panic(err)
	}

	query := `
	CREATE TABLE IF NOT EXISTS items (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		value TEXT
	);`
	_, err = db.Exec(query)
	if err != nil {
		panic(err)
	}
}

func getItems(c echo.Context) error {
	rows, err := db.Query("SELECT id, name, value FROM items")
	if err != nil {
		return err
	}
	defer rows.Close()

	items := []Item{}
	for rows.Next() {
		var item Item
		if err := rows.Scan(&item.ID, &item.Name, &item.Value); err != nil {
			return err
		}
		items = append(items, item)
	}

	return c.JSON(http.StatusOK, items)
}

func createItem(c echo.Context) error {
	item := new(Item)
	if err := c.Bind(item); err != nil {
		return err
	}

	stmt, err := db.Prepare("INSERT INTO items (name, value) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(item.Name, item.Value)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, item)
}

func updateItem(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	item := new(Item)
	if err := c.Bind(item); err != nil {
		return err
	}

	stmt, err := db.Prepare("UPDATE items SET name = ?, value = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(item.Name, item.Value, id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, item)
}

func deleteItem(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	stmt, err := db.Prepare("DELETE FROM items WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func main() {
	initDB()
	e := echo.New()

	e.GET("/items", getItems)
	e.POST("/items", createItem)
	e.PUT("/items/:id", updateItem)
	e.DELETE("/items/:id", deleteItem)

	e.Start(":8080")
}
