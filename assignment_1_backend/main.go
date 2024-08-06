package main

import (
    "database/sql"
    "encoding/json"
    "log"
    "net/http"

    _ "github.com/mattn/go-sqlite3"
    "github.com/gorilla/mux"
)

var db *sql.DB

func main() {
    var err error

    // Koneksi ke database SQLite
    // db, err = sql.Open("sqlite3", "./tugaspertama.db")
    db, err = sql.Open("sqlite3", "C:\\sqlite\\database\\akhirnya.db")


    if err != nil {
        log.Fatalf("Error opening database: %v", err)
    }
    defer db.Close()

    // Cek koneksi ke database
    if err = db.Ping(); err != nil {
        log.Fatalf("Error connecting to database: %v", err)
    }
    log.Println("Database connected successfully")

    // Buat tabel jika belum ada
    if err = createTable(); err != nil {
        log.Fatalf("Error creating table: %v", err)
    }

    // Setup router
    r := mux.NewRouter()
    r.HandleFunc("/items", getItems).Methods("GET")
    r.HandleFunc("/items", createItems).Methods("POST")
    r.HandleFunc("/items/{id}", updateItem).Methods("PUT")
    r.HandleFunc("/items/{id}", deleteItem).Methods("DELETE")

    log.Fatal(http.ListenAndServe(":8000", r))
}

func createTable() error {
    createTableSQL := `CREATE TABLE IF NOT EXISTS items (
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        name TEXT
    );`
    _, err := db.Exec(createTableSQL)
    if err != nil {
        return err
    }
    log.Println("Table 'items' checked/created successfully")
    return nil
}

func getItems(w http.ResponseWriter, r *http.Request) {
    rows, err := db.Query("SELECT id, name FROM items")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var items []map[string]interface{}
    for rows.Next() {
        var id int
        var name string
        if err := rows.Scan(&id, &name); err != nil {
            log.Printf("Failed to scan row: %v", err)
            continue
        }
        items = append(items, map[string]interface{}{"id": id, "name": name})
    }

    json.NewEncoder(w).Encode(items)
}

func createItems(w http.ResponseWriter, r *http.Request) {
    var item map[string]interface{}
    if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    _, err := db.Exec("INSERT INTO items (name) VALUES (?)", item["name"])
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(item)
}

func updateItem(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var item map[string]interface{}
    if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    _, err := db.Exec("UPDATE items SET name = ? WHERE id = ?", item["name"], params["id"])
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(item)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)

    _, err := db.Exec("DELETE FROM items WHERE id = ?", params["id"])
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
