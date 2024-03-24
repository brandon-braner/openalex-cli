package database

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"io"
	"log"
	"net/http"
	"os"
)

func Connect() (*sql.DB, error) {
	// get the database connection string from the environment
	connectionString := os.Getenv("OPEN_ALEX_CLI_DB")
	if connectionString == "" {
		return nil, errors.New("OPEN_ALEX_CLI_DB environment variable not set")
	}
	// Connect to the database
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	// Check if the connection is successful
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Successfully connected to the database")
	return db, nil

}

func RunRemoteSQL(url string) {
	// only pass a url into this that you have control over to prevent random sql to be ran
	db, err := Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error getting sql file at url %s", url)
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
}
