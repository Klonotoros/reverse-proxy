package database

import (
	"database/sql"
	_ "github.com/ClickHouse/clickhouse-go/v2"
)

func InitDB() *sql.DB {
	DB, err := sql.Open("clickhouse", "tcp://localhost:9000?username=userTest&password=passwordTest")
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
