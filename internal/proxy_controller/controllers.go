package proxy_controller

import (
	"github.com/gin-gonic/gin"
	"proxy-server/internal/service"
)

type Controllers interface {
	Proxy() ProxyController
	Route(server *gin.Engine)
}

type controllers struct {
	proxyController ProxyController
}

func NewControllers(services service.Services) Controllers {
	return &controllers{
		proxyController: newProxyController(services.Proxy()),
	}
}

func (c *controllers) Route(server *gin.Engine) {
	server.Any("/*url", c.proxyController.HandleRequest)
}

func (c *controllers) Proxy() ProxyController {
	return c.proxyController
}
