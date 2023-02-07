package main

import (
	"log"

	// Import interno de packages do próprio sistema
	"microservice_user/config"
	"microservice_user/pkg/database"
	"microservice_user/pkg/routes"
	"microservice_user/pkg/server"
	"microservice_user/pkg/service"
)

// Função principal (primeira executada) - chama config para fazer conexão BD, service, server, router e roda servidor http
func main() {
	// Atribui o endereço da estrutura de uma configuração padrão do sistema
	conf := &config.Config{}

	// Atribui para conf as novas configurações do sistema
	conf = config.NewConfig()

	// Pega pool de conexão do Database (config passadas anteriorente pelo Json para Database)
	dbpool := database.NewDB(conf)
	// Se criou uma conexão com as configurações passadas para o Database - Imprima essa mensagem de sucesso
	if dbpool != nil {
		log.Print("Successfully connected")
	}

	// Cria serviços de um produto (CRUD) com a pool de conexão passada por parâmetro
	service := service.NewUserService(dbpool)

	// Cria servidor HTTP com as config passadas por parâmetro
	serv := server.NewServer(conf)

	// Cria rotas passsando o servidor HTTP e os serviços do produto (CRUD)
	router := routes.ConfigRoutes(serv.SERVER, service)

	// Coloca servidor para rodar passando as rotas, servidor HTTP e serviços do produto (CRUD) como parâmetro
	server.Run(router, serv, service)
}
