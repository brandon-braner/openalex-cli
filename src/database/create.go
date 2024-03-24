package database

import (
	"fmt"
)

func CreateTables() {
	sqlUrl := "https://raw.githubusercontent.com/brandon-braner/openalex_sql/main/create_table.sql"
	RunRemoteSQL(sqlUrl)
	fmt.Println("Tables created successfully")
}

func CreateIndexes() {
	sqlUrl := "https://raw.githubusercontent.com/brandon-braner/openalex_sql/main/create_indexes.sql"
	RunRemoteSQL(sqlUrl)
	fmt.Println("Indexes created successfully")
}
