package database

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func CreateTables() {
	db, err := Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlUrl := "https://raw.githubusercontent.com/brandon-braner/openalex_sql/main/create_table.sql"

	resp, err := http.Get(sqlUrl)
	if err != nil {
		fmt.Println("Error getting sql file at url %s", sqlUrl)
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	query := string(body)

	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Tables created successfully")
}

func CreateIndexes() {
	db, err := Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlUrl := "https://raw.githubusercontent.com/brandon-braner/openalex_sql/main/create_indexes.sql"

	resp, err := http.Get(sqlUrl)
	if err != nil {
		fmt.Println("Error getting sql file at url %s", sqlUrl)
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	query := string(body)

	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Indexes created successfully")
}
