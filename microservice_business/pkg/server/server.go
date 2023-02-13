package server

import (
	"log"

	// Import interno de packages do próprio sistema
	"microservice_business/config"
	"microservice_business/pkg/service"

	// Import externo do github
	"github.com/gin-gonic/gin"
)

type Server struct {
	// Porta do servidor
	SRV_PORT string

	SERVER *gin.Engine
}

func NewServer(conf *config.Config) Server {
	return Server{
		SRV_PORT: conf.SRV_PORT,
		SERVER:   gin.Default(),
	}
}

func Run(router *gin.Engine, server Server, service service.BusinessServiceInterface) {
	log.Print("Server is running at port: ", server.SRV_PORT)

	log.Fatal(router.Run(":" + server.SRV_PORT))
}
