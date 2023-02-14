package server

import (
	"log"

	"microservice_subject/config"
	"microservice_subject/pkg/service"

	// Import interno de packages do próprio sistema

	// Import externo do github
	"github.com/gin-gonic/gin"
)

// Estrutura de dados para armazenar o servidor HTTP
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

func Run(router *gin.Engine, server Server, service service.SubjectServiceInterface) {
	log.Print("Server is running at port: ", server.SRV_PORT)

	log.Fatal(router.Run(":" + server.SRV_PORT))
}
