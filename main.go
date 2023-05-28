package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type book struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func getBooks(db *sql.DB) (books []book, err error) {
	var book book

	rows, err := db.Query("SELECT * FROM local_db.books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
		); err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}

func getBooksHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		books, err := getBooks(db)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{
				"error": err.Error(),
			})

			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(books)
	}
}

func main() {
	db, err := sql.Open("mysql", "root:pass@tcp(localhost)/local_db?parseTime=true")
	if err != nil {
		log.Fatalf("failed to open db connection: %v\n", err)
	}

	http.HandleFunc("/books", getBooksHandler(db))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("failed to start http server: %v\n", err)
	}
}
