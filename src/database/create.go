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

	sqlUrl := "https://gist.githubusercontent.com/brandon-braner/38d085d119394889d1f4ec55e631f6ce/raw/349a25414bae30843b1f50b0074f8c366564bf6b/create_tables.sql"

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
