package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func Connect() (*sql.DB, error) {
	// Connect to the database
	connectionString := "postgres://postgres.rnhjsxpzbgmdkwgypkga:PeanutButterCup1!@aws-0-us-east-1.pooler.supabase.com:5432/postgres"
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
