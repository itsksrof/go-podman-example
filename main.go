package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Channel struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func main() {
	db, err := sql.Open("mysql", "root:pass@tcp(localhost)/local_db?parseTime=true")
	if err != nil {
		log.Fatalf("failed to open db connection: %v\n", err)
	}

	http.HandleFunc("/channels", func(w http.ResponseWriter, r *http.Request) {
		channels, err := getChannels(db)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
				return
			}

			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(channels)
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("failed to start http server: %v\n", err)
	}
}

func getChannels(db *sql.DB) ([]Channel, error) {
	rows, err := db.Query("SELECT * FROM local_db.channels")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	channels := []Channel{}
	channel := Channel{}

	for rows.Next() {
		if err := rows.Scan(&channel.ID, &channel.Name); err != nil {
			return nil, err
		}

		channels = append(channels, channel)
	}

	return channels, nil
}
