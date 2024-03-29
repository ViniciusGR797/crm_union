package server

import (
	"log"

	// Import interno de packages do próprio sistema
	"microservice_spreadsheet/config"
	"microservice_spreadsheet/pkg/service"

	// Import externo do github
	"github.com/gin-gonic/gin"
)

// Estrutura de dados para armazenar o servidor HTTP
type Server struct {
	// Porta do servidor
	SRV_PORT string

	// Ponteiro de servidor do framework gin
	SERVER *gin.Engine
}

// NewServer Cria novo servidor HTTP, de acordo com as config passadas por parâmetro
func NewServer(conf *config.Config) Server {
	return Server{
		SRV_PORT: conf.SRV_PORT,
		SERVER:   gin.Default(),
	}
}

// Run Rodar servidor HTTP, tendo as rotas do framework gin, servidor HTTP, serviço CRUD de produto
func Run(router *gin.Engine, server Server, service service.RemarkServiceInterface) {
	// Imprime que servidor HTTP está rodando na porta tal
	log.Print("Server is running at port: ", server.SRV_PORT)

	// Roda servidor HTTP com as rotas e a porta do servidor passadas por parâmetro (caso ser erro dá Fatal erro - fecha o sistema)
	log.Fatal(router.Run(":" + server.SRV_PORT))
}
