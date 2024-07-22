# Echo SQLite Example

This project is a simple implementation of a REST API using Go, Echo, and SQLite. It demonstrates how to create CRUD (Create, Read, Update, Delete) endpoints to manage items in an SQLite database.

## Prerequisites

Before you begin, ensure you have the following installed:

- [Go](https://golang.org/dl/)
- [TDM-GCC](https://jmeubank.github.io/tdm-gcc/) or MinGW for C compiler on Windows.
- [Postman](https://www.postman.com/downloads/) for API testing.

## Installation

Install my-project with npm

1. Clone this repository or create a new project directory:

   ```sh
   mkdir echo-sqlite-example
   cd echo-sqlite-example
   ```

2. Initialize the Go module:

   ```sh
   go mod init echo-sqlite-example
   ```

3. Install dependencies:

   ```sh
   go get github.com/labstack/echo/v4
   go get github.com/mattn/go-sqlite3
   ```

4. Create a `main.go` file with the following content:

   ```go
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
       db, err = sql.Open("sqlite3", "./test.db")
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
   ```

## Running the Application

1. Open a terminal or command prompt.
2. Enable CGO:

   ```sh
   set CGO_ENABLED=1
   ```

3. Run the application:

   ```sh
   go run main.go
   ```

The application will run on `http://localhost:8080`.

## Accessing Endpoints with Postman

### GET /items

- Method: GET
- URL: `http://localhost:8080/items`
- Description: Retrieves a list of all items.

### POST /items

- Method: POST
- URL: `http://localhost:8080/items`
- Body:
  ```json
  {
    "name": "ItemName",
    "value": "ItemValue"
  }
  ```
- Description: Adds a new item.

### PUT /items/:id

- Method: PUT
- URL: `http://localhost:8080/items/{id}`
- Body:
  ```json
  {
    "name": "UpdatedItemName",
    "value": "UpdatedItemValue"
  }
  ```
- Description: Updates an item by ID.

### DELETE /items/:id

- Method: DELETE
- URL: `http://localhost:8080/items/{id}`
- Description: Deletes an item by ID.
