package server

import (
	"log"

	// Import interno de packages do próprio sistema
	"microservice_group/config"
<<<<<<< HEAD
=======

>>>>>>> c159f9d2de112426f41a26075473b24bb06e931f
	"microservice_group/pkg/service"

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


func Run(router *gin.Engine, server Server, service service.GroupServiceInterface) {
	log.Print("Server is running at port: ", server.SRV_PORT)
	
	log.Fatal(router.Run(":" + server.SRV_PORT))
}
