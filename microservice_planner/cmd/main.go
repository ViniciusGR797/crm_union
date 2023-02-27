package main

import (
	"encoding/json"
	"io"
	"log"
	"os"

	// Import interno de packages do próprio sistema
	"microservice_planner/config"
	"microservice_planner/pkg/database"
	"microservice_planner/pkg/routes"
	"microservice_planner/pkg/security"
	"microservice_planner/pkg/server"
	"microservice_planner/pkg/service"
)

// Função principal (primeira executada) - chama config para fazer conexão BD, service, server, router e roda servidor http
func main() {
	// Atribui o endereço da estrutura de uma configuração padrão do sistema
	default_conf := &config.Config{}

	// Abre o arquivo JSON com as variáveis de ambiente
	file, err := os.Open("./env.json") // file.json has the json content
	if err != nil {
		panic(err)
	}

	// Lé todo JSON e transforma em um JSON byte
	jsonByte, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	// Converte JSON byte em uma struct, no caso a struct default_conf
	if err := json.Unmarshal(jsonByte, &default_conf); err != nil {
		panic(err)
	}

	// Atribui para conf as novas configurações do sistema
	conf := config.NewConfig(default_conf)

	// Pega pool de conexão do Database (config passadas anteriorente pelo Json para Database)
	dbpool := database.NewDB(conf)
	// Se criou uma conexão com as configurações passadas para o Database - Imprima essa mensagem de sucesso
	if dbpool != nil {
		log.Print("Successfully connected")
	}

	// Cria serviços de um planner (CRUD) com a pool de conexão passada por parâmetro
	service := service.NewPlannerService(dbpool)

	// Configura a chave de segurança dos tokens
	err = security.SecretConfig(conf)
	if err != nil {
		panic(err)
	}

	// Cria servidor HTTP com as config passadas por parâmetro
	serv := server.NewServer(conf)

	// Cria rotas passsando o servidor HTTP e os serviços do planner (CRUD)
	router := routes.ConfigRoutes(serv.SERVER, service)

	// Se tiver ativada a interface de usuário, criar as rotas para o front end (WEB UI)
	// if conf.WEB_UI {
	// 	webui.RegisterUIHandlers(router)
	// }

	// Coloca servidor para rodar passando as rotas, servidor HTTP e serviços do planner (CRUD) como parâmetro
	server.Run(router, serv, service)
}
