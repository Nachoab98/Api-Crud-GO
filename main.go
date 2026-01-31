package main

import (
	"database/sql"
	"log"
	"net/http"
	"goApi/internal/service"
	"goApi/internal/store"
	"goApi/internal/transport"
)

func main() {
	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		log.Fatal(err)
}
	defer db.Close()
	q:= `CREATE TABLE IF NOT EXISTS books (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		author TEXT NOT NULL
	)`
	if _, err = db.Exec(q); err != nil {
		log.Fatal(err.Error())
	}
	bookStore := store.NewStore(db)
	bookService := service.New(bookStore)
	bookHandler := transport.New(bookService)
	http.HandleFunc("/books", bookHandler.HandleBooks)
	http.HandleFunc("/books/", bookHandler.HandleBookByID)

	log.Fatal(http.ListenAndServe(":8080", nil))

}