package server

import (
	"log"

	// Import interno de packages do próprio sistema
	"microservice_user/config"
	"microservice_user/pkg/service"

	// Import externo do github
	"github.com/gin-gonic/gin"
)

// Server é Estrutura de dados para armazenar o servidor HTTP
type Server struct {
	// Porta do servidor
	SRV_PORT string

	// Ponteiro de servidor do framework gin
	SERVER *gin.Engine
}

// NewServer cria novo servidor HTTP, de acordo com as config passadas por parâmetro
func NewServer(conf *config.Config) Server {
	return Server{
		SRV_PORT: conf.SRV_PORT,
		SERVER:   gin.Default(),
	}
}

// Run roda o servidor HTTP, tendo as rotas do framework gin, servidor HTTP, serviço CRUD de user
func Run(router *gin.Engine, server Server, service service.UserServiceInterface) {
	// Imprime que servidor HTTP está rodando na porta tal
	log.Print("Server is running at port: ", server.SRV_PORT)

	// Roda servidor HTTP com as rotas e a porta do servidor passadas por parâmetro (caso ser erro dá Fatal erro - fecha o sistema)
	log.Fatal(router.Run(":" + server.SRV_PORT))
}
