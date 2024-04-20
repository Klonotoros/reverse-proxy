package server_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"proxy-server/internal/dto"
	"proxy-server/internal/service"
)

type RecordController interface {
	DownloadRecords(*gin.Context)
	GetNumberOfRecords(*gin.Context)
}

type serverController struct {
	recordService service.RecordService
}

func newServerController(recordService service.RecordService) RecordController {
	return &serverController{recordService: recordService}
}

func (s *serverController) GetNumberOfRecords(ctx *gin.Context) {
	numberOfRecords, err := s.recordService.CountRecords()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	totalResponse := dto.RecordNumberResponse{Total: numberOfRecords}
	ctx.JSON(http.StatusOK, totalResponse)
}

func (s *serverController) DownloadRecords(ctx *gin.Context) {
	csvData, err := s.recordService.GenerateCSV()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.Header("Content-Type", "text/csv")
	ctx.Header("Content-Disposition", "attachment; filename=records.csv")

	ctx.Data(http.StatusOK, "text/csv", csvData)
}
