package service

import (
	"proxy-server/internal/model"
	"proxy-server/internal/repository"
	"time"
)

type ProxyService interface {
	HandleRequest(method string, url string, clientIP string) error
}

type proxyService struct {
	recordRepository repository.RecordRepository
}

func newProxyService(recordRepository repository.RecordRepository) ProxyService {
	return &proxyService{recordRepository: recordRepository}
}

func (p *proxyService) HandleRequest(method string, url string, clientIP string) error {
	return p.recordRepository.Save(model.Record{
		Timestamp: time.Now().Unix(),
		URL:       url,
		Method:    method,
		IP:        clientIP,
	})
}
