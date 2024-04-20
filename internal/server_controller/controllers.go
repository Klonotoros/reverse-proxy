package server_controller

import (
	"github.com/gin-gonic/gin"
	"proxy-server/internal/service"
)

type Controllers interface {
	Server() RecordController
	Route(server *gin.Engine)
}

type controllers struct {
	recordController RecordController
}

func NewControllers(services service.Services) Controllers {
	return &controllers{
		recordController: newServerController(services.Record()),
	}
}

func (c *controllers) Route(server *gin.Engine) {
	server.GET("/download", c.recordController.DownloadRecords)
	server.GET("/info", c.recordController.GetNumberOfRecords)
}

func (c *controllers) Server() RecordController {
	return c.recordController
}
