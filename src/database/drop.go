package database

import (
	"fmt"
)

func DropTables() {
	sqlUrl := "https://raw.githubusercontent.com/brandon-braner/openalex_sql/main/drop_all_tables.sql"
	RunRemoteSQL(sqlUrl)
	fmt.Println("Tables dropped successfully")
}

func DropIndexes() {
	sqlUrl := "https://raw.githubusercontent.com/brandon-braner/openalex_sql/main/drop_indexes.sql"
	RunRemoteSQL(sqlUrl)
	fmt.Println("Indexes dropped successfully")
}
