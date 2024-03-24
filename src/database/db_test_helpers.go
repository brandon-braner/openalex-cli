package database

import (
	"database/sql"
	"log"
)

func SetupTestDb() *sql.DB {
	db, err := Connect()
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("DROP SCHEMA if exists openalex CASCADE")
	if err != nil {
		log.Fatal(err)

	}
	_, err = db.Exec("CREATE SCHEMA if not exists openalex")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func TeardownTestDb() {
	db, err := Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("DROP SCHEMA if exists openalex CASCADE")

	if err != nil {
		log.Fatal(err)

	}
}
