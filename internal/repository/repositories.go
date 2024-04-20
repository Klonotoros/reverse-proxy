package repository

import "database/sql"

type Repositories interface {
	Record() RecordRepository
}

type repositories struct {
	recordRepository RecordRepository
}

func NewRepositories(db *sql.DB) Repositories {
	return &repositories{
		recordRepository: newRecordRepository(db),
	}
}

func (r repositories) Record() RecordRepository {
	return r.recordRepository
}
