package service

import (
	"bytes"
	"encoding/csv"
	"proxy-server/internal/model"
	"proxy-server/internal/repository"
	"strconv"
)

type RecordService interface {
	GenerateCSV() ([]byte, error)
	CountRecords() (int, error)
	SaveRecord(record model.Record) error
}

type recordService struct {
	recordRepository repository.RecordRepository
}

func newRecordService(recordRepository repository.RecordRepository) RecordService {
	return &recordService{recordRepository: recordRepository}
}

func (r *recordService) GenerateCSV() ([]byte, error) {

	records, err := r.recordRepository.GetRecords()
	if err != nil {
		return nil, err
	}

	var csvData bytes.Buffer
	csvWriter := csv.NewWriter(&csvData)

	for _, record := range records {
		err := csvWriter.Write([]string{strconv.FormatInt(record.Timestamp, 10), record.URL, record.Method, record.IP})
		if err != nil {
			return nil, err
		}
	}
	csvWriter.Flush()

	return csvData.Bytes(), nil
}

func (r *recordService) CountRecords() (int, error) {
	return r.recordRepository.CountRecords()
}

func (r *recordService) SaveRecord(record model.Record) error {
	return r.recordRepository.Save(record)
}
