package proxy_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"net/url"
	"proxy-server/internal/service"
	"strings"
)

type ProxyController interface {
	HandleRequest(ctx *gin.Context)
}

type proxyController struct {
	proxyService service.ProxyService
}

func newProxyController(proxyService service.ProxyService) ProxyController {
	return &proxyController{proxyService: proxyService}
}

func (p *proxyController) HandleRequest(ctx *gin.Context) {
	urlParam := strings.TrimPrefix(ctx.Param("url"), "/")

	if !strings.HasPrefix(urlParam, "http://") && !strings.HasPrefix(urlParam, "https://") {
		urlParam = "http://" + urlParam
	}

	targetURL, err := url.Parse(urlParam)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	proxy.Director = func(req *http.Request) {
		req.Header = ctx.Request.Header
		req.Host = targetURL.Host
		req.URL.Scheme = targetURL.Scheme
		req.URL.Host = targetURL.Host
		req.URL.Path = ""
	}

	proxy.ModifyResponse = func(resp *http.Response) error {
		return p.proxyService.HandleRequest(ctx.Request.Method, urlParam, ctx.ClientIP())
	}

	proxy.ServeHTTP(ctx.Writer, ctx.Request)
}
