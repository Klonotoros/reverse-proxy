package main

import (
	_ "github.com/ClickHouse/clickhouse-go/v2"
	"github.com/gin-gonic/gin"
	"proxy-server/internal/database"
	"proxy-server/internal/proxy_controller"
	"proxy-server/internal/repository"
	"proxy-server/internal/server_controller"
	"proxy-server/internal/service"
)

func main() {

	db := database.InitDB()
	repositories := repository.NewRepositories(db)
	services := service.NewServices(repositories)

	proxy := gin.Default()
	server := gin.Default()
	proxyController := proxy_controller.NewControllers(services)
	proxyController.Route(proxy)
	go func() {
		err := proxy.Run(":3101")
		if err != nil {
			panic(err)
		}

	}()

	serverController := server_controller.NewControllers(services)
	serverController.Route(server)
	err := server.Run(":3102")
	if err != nil {
		panic(err)
	}
}
