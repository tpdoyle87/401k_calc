package main

import (
	"401k_calculator/api"
	_ "401k_calculator/api"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	db, err := sql.Open("postgres", "user=postgres dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	apiHandler := api.SetupAPI()
	log.Fatal(http.ListenAndServe(":8080", apiHandler))
}
