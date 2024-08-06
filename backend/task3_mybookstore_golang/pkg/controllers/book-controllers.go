package controllers

import (
    "encoding/json"
    "fmt"
    "task3_mybookstore_golang/pkg/models"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
    books := models.GetAllBooks()
    res, _ := json.Marshal(books)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    Id, err := strconv.ParseInt(vars["bookId"], 0, 0)
    if err != nil {
        fmt.Println("error while parsing")
    }
    bookDetails, _ := models.GetBookById(Id)
    res, _ := json.Marshal(bookDetails)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
    CreateBook := &models.Book{}
    err := json.NewDecoder(r.Body).Decode(CreateBook)
    if err != nil {
        fmt.Println("error while decoding")
    }
    b := CreateBook.CreateBook()
    res, _ := json.Marshal(b)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    Id, err := strconv.ParseInt(vars["bookId"], 0, 0)
    if err != nil {
        fmt.Println("error while parsing")
    }
    book := models.DeleteBook(Id)
    res, _ := json.Marshal(book)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
    var updateBook = &models.Book{}
    err := json.NewDecoder(r.Body).Decode(updateBook)
    if err != nil {
        fmt.Println("error while decoding")
    }
    vars := mux.Vars(r)
    Id, err := strconv.ParseInt(vars["bookId"], 0, 0)
    if err != nil {
        fmt.Println("error while parsing")
    }
    bookDetails, db := models.GetBookById(Id)
    if updateBook.Title != "" {
        bookDetails.Title = updateBook.Title
    }
    if updateBook.Author != "" {
        bookDetails.Author = updateBook.Author
    }
    if updateBook.Publisher != "" {
        bookDetails.Publisher = updateBook.Publisher
    }
    db.Save(&bookDetails)
    res, _ := json.Marshal(bookDetails)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}
