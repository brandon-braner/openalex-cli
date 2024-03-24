package database

import (
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"log"
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
