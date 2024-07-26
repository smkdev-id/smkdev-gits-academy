package handlers

import (
	"database/sql"
	"net/http"
	"task-2-golang-todo-crud-api/models"

	"github.com/labstack/echo/v4"
)

// GetItems mengembalikan daftar semua item
func GetItems(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		rows, err := db.Query("SELECT id, title, description FROM items")
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		defer rows.Close()

		var items []models.Item
		for rows.Next() {
			var item models.Item
			if err := rows.Scan(&item.ID, &item.Title, &item.Description); err != nil {
				return c.JSON(http.StatusInternalServerError, err)
			}
			items = append(items, item)
		}

		return c.JSON(http.StatusOK, items)
	}
}

// GetItem mengembalikan item berdasarkan ID
func GetItem(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		var item models.Item
		err := db.QueryRow("SELECT id, title, description FROM items WHERE id = ?", id).Scan(&item.ID, &item.Title, &item.Description)
		if err != nil {
			if err == sql.ErrNoRows {
				return c.JSON(http.StatusNotFound, map[string]string{
					"message": "Item not found",
				})
			}
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": "Failed to retrieve item",
			})
		}

		return c.JSON(http.StatusOK, item)
	}
}

// CreateItem menambahkan item baru
func CreateItem(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var item models.Item
		if err := c.Bind(&item); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": "Invalid request payload",
			})
		}

		result, err := db.Exec("INSERT INTO items (title, description) VALUES (?, ?)", item.Title, item.Description)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": "Failed to create item. Erorr: err",
			})
		}

		id, _ := result.LastInsertId()
		item.ID = int(id)

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "Item created successfully",
			"item":    item,
		})
	}
}

// UpdateItem memperbarui item berdasarkan ID
func UpdateItem(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		var item models.Item
		if err := c.Bind(&item); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": "Invalid request payload",
			})
		}

		// Check if the item exists
		var existingItem models.Item
		err := db.QueryRow("SELECT id, title, description FROM items WHERE id = ?", id).Scan(&existingItem.ID, &existingItem.Title, &existingItem.Description)
		if err != nil {
			if err == sql.ErrNoRows {
				return c.JSON(http.StatusNotFound, map[string]string{
					"message": "Item not found",
				})
			}
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": "Failed to retrieve item",
			})
		}

		_, err = db.Exec("UPDATE items SET title = ?, description = ? WHERE id = ?", item.Title, item.Description, id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": "Failed to update item",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Item updated successfully",
			"item": map[string]string{
				"title":       item.Title,
				"description": item.Description,
			},
		})
	}
}

// DeleteItem menghapus item berdasarkan ID
func DeleteItem(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		_, err := db.Exec("DELETE FROM items WHERE id = ?", id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": "Failed to delete item",
			})
		}

		return c.JSON(http.StatusOK, map[string]string{
			"message": "Item deleted successfully",
		})
	}
}
