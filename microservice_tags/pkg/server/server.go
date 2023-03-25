package server

import (
	"log"

	// Import interno de packages do próprio sistema
	"microservice_tags/config"
	"microservice_tags/pkg/service"

	// Import externo do github
	"github.com/gin-gonic/gin"
)

// Server estrutura de dados para Server
type Server struct {
	// Porta do servidor
	SRV_PORT string

	SERVER *gin.Engine
}

// NewServer cria um novo Server
func NewServer(conf *config.Config) Server {
	return Server{
		SRV_PORT: conf.SRV_PORT,
		SERVER:   gin.Default(),
	}
}

// Run sobe um Server
func Run(router *gin.Engine, server Server, service service.TagsServiceInterface) {
	log.Print("Server is running at port: ", server.SRV_PORT)

	log.Fatal(router.Run(":" + server.SRV_PORT))
}
