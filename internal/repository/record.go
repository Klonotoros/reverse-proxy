package repository

import (
	"database/sql"
	"log"
	"proxy-server/internal/model"
)

type RecordRepository interface {
	CountRecords() (int, error)
	GetRecords() ([]model.Record, error)
	Save(model.Record) error
}

type record struct {
	db *sql.DB
}

func newRecordRepository(db *sql.DB) RecordRepository {
	return &record{db: db}
}

func (r *record) CountRecords() (int, error) {
	var count int
	err := r.db.QueryRow("SELECT COUNT(*) FROM record").Scan(&count)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return count, nil
}

func (r *record) GetRecords() ([]model.Record, error) {
	rows, err := r.db.Query("SELECT * FROM record ORDER BY Timestamp ")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println(err)
		}
	}(rows)

	var records []model.Record
	for rows.Next() {
		var record model.Record
		err = rows.Scan(&record.Timestamp, &record.URL, &record.Method, &record.IP)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		records = append(records, record)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return records, nil
}

func (r *record) Save(record model.Record) error {
	_, err := r.db.Exec("INSERT INTO record(Timestamp, URL, Method, IP) VALUES (?, ?, ?, ?)", record.Timestamp, record.URL, record.Method, record.IP)
	return err
}
