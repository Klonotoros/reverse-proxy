package service

import "proxy-server/internal/repository"

type Services interface {
	Record() RecordService
	Proxy() ProxyService
}

type services struct {
	recordService RecordService
	proxyService  ProxyService
}

func NewServices(repositories repository.Repositories) Services {
	return &services{
		recordService: newRecordService(repositories.Record()),
		proxyService:  newProxyService(repositories.Record()),
	}
}

func (s *services) Record() RecordService {
	return s.recordService
}

func (s *services) Proxy() ProxyService {
	return s.proxyService
}
