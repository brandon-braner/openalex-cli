package database

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func DropTables() {
	db, err := Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlUrl := "https://raw.githubusercontent.com/brandon-braner/openalex_sql/main/drop_all_tables.sql"

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
	fmt.Println("Tables dropped successfully")
}

func DropIndexes() {
	db, err := Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlUrl := "https://raw.githubusercontent.com/brandon-braner/openalex_sql/main/drop_indexes.sql"

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
	fmt.Println("Indexes dropped successfully")
}
