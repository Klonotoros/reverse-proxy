package database

import (
	"database/sql"
	_ "github.com/ClickHouse/clickhouse-go/v2"
	"os"
)

func InitDB() *sql.DB {
	dsn := os.Getenv("CLICKHOUSE_DSN")
	DB, err := sql.Open("clickhouse", dsn)
	if err != nil {
		panic("Database could not connect: " + err.Error())
		return nil
	}

	err = createTables(DB)
	if err != nil {
		panic("Could not create tables: " + err.Error())
		return nil
	}

	return DB
}

func createTables(db *sql.DB) error {
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS record(
			Timestamp Int64 PRIMARY KEY,
			URL String,
			Method String,
			IP String
		) ENGINE = MergeTree()
		ORDER BY Timestamp
	`
	_, err := db.Exec(createTableQuery)
	if err != nil {
		panic("Could not create records table.")
	}

	return err
}
